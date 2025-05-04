package storage

import (
	"context"
	"fmt"
	"src/common/setting"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"

	"mime/multipart"
	"src/common/ctype"
	"src/util/errutil"
	"src/util/i18nmsg"
	"time"
)

type adapter struct {
	client     *s3.Client
	bucketName string
}

type FileInfo struct {
	FileName string
	FileType string
	FileURL  string
	FileSize int
}

type FileInfoMap map[string]FileInfo

func New() adapter {
	region := setting.S3_REGION()
	cfg, _ := config.LoadDefaultConfig(
		context.TODO(),
		config.WithCredentialsProvider(
			credentials.NewStaticCredentialsProvider(
				setting.S3_ACCESS_KEY_ID(), setting.S3_SECRET_ACCESS_KEY(), "",
			)),
		config.WithRegion(region),
	)
	s3Client := s3.NewFromConfig(cfg, func(o *s3.Options) {
		o.BaseEndpoint = aws.String(
			fmt.Sprintf(
				"https://%s.r2.cloudflarestorage.com",
				setting.S3_ACCOUNT_ID(),
			),
		)
	})

	bucketName := setting.S3_BUCKET_NAME()
	return adapter{client: s3Client, bucketName: bucketName}
}

func (a *adapter) Uploads(
	ctx context.Context,
	fileHeaderMap map[string][]*multipart.FileHeader,
	folder string,
) (FileInfoMap, error) {
	result := FileInfoMap{}
	errChan := make(chan error, len(fileHeaderMap))
	doneChan := make(chan struct{}, len(fileHeaderMap))

	for fieldName, fileHeaderMap := range fileHeaderMap {
		go func(fieldName string, fileHeaderMap []*multipart.FileHeader) {
			for _, fileHeader := range fileHeaderMap {
				fileInfo, err := a.upload(ctx, fileHeader, folder)
				if err != nil {
					errChan <- err
					return
				}
				result[fieldName] = fileInfo
			}
			doneChan <- struct{}{}
		}(fieldName, fileHeaderMap)
	}

	for i := 0; i < len(fileHeaderMap); i++ {
		select {
		case err := <-errChan:
			return nil, errutil.NewRaw(err.Error())
		case <-doneChan:
		}
	}

	return result, nil
}

func (a *adapter) upload(
	ctx context.Context,
	fileHeader *multipart.FileHeader,
	folder string,
) (FileInfo, error) {
	emptyResult := FileInfo{}
	file, err := fileHeader.Open()
	if err != nil {
		return emptyResult, errutil.NewWithArgs(
			i18nmsg.FailedToOpenFile,
			ctype.Dict{
				"Filename": fileHeader.Filename,
			},
		)
	}
	defer file.Close()

	// Unique key for each file
	key := fmt.Sprintf("%s/%d_%s", folder, time.Now().Unix(), fileHeader.Filename)
	_, err = a.client.PutObject(ctx, &s3.PutObjectInput{
		Bucket: aws.String(a.bucketName),
		Key:    aws.String(key),
		Body:   file,
		ACL:    "public-read",
	})
	if err != nil {
		fmt.Println(err)
		return emptyResult, errutil.New(i18nmsg.FailedToUploadFileToS3)
	}

	// Generate the S3 URL
	fileUrl := fmt.Sprintf("%s/%s", setting.S3_ENDPOINT_URL(), key)
	fileName := fileHeader.Filename
	fileType := fileHeader.Header.Get("Content-Type")
	fileSize := int(fileHeader.Size)
	return FileInfo{
		FileName: fileName,
		FileType: fileType,
		FileURL:  fileUrl,
		FileSize: fileSize,
	}, nil
}
