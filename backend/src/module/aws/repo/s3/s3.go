package s3

import (
	"context"
	"fmt"
	"mime/multipart"
	"src/common/ctype"
	"src/common/setting"
	"src/util/errutil"
	"src/util/localeutil"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

type Repo struct {
	client     *s3.Client
	bucketName string
	region     string
}

func New(client *s3.Client) Repo {
	bucketName := setting.S3_BUCKET_NAME
	region := setting.S3_REGION
	return Repo{client: client, bucketName: bucketName, region: region}
}

func (u *Repo) Upload(
	ctx context.Context,
	folder string,
	fileHeader *multipart.FileHeader,
) (string, error) {
	localizer := localeutil.Get()
	file, err := fileHeader.Open()
	if err != nil {
		msg := localizer.MustLocalize(&i18n.LocalizeConfig{
			DefaultMessage: localeutil.FailedToOpenFile,
			TemplateData: ctype.Dict{
				"Filename": fileHeader.Filename,
			},
		})
		return "", errutil.New("", []string{msg})

	}
	defer file.Close()

	// Unique key for each file
	key := fmt.Sprintf("%s/%d_%s", folder, time.Now().Unix(), fileHeader.Filename)
	_, err = u.client.PutObject(ctx, &s3.PutObjectInput{
		Bucket: aws.String(u.bucketName),
		Key:    aws.String(key),
		Body:   file,
		ACL:    "public-read",
	})
	if err != nil {
		fmt.Println(err)
		msg := localizer.MustLocalize(&i18n.LocalizeConfig{
			DefaultMessage: localeutil.FailedToUploadFileToS3,
		})
		return "", errutil.New("", []string{msg})
	}

	// Generate the S3 URL
	s3URL := fmt.Sprintf("%s/%s", setting.S3_ENDPOINT_URL, key)
	return s3URL, nil
}

// write Uploads function that receive map of fileHeaders, upload to S3 parallelly using goroutine return map of fieldName and s3URL, reuse Upload function
func (u *Repo) Uploads(
	ctx context.Context,
	folder string,
	files map[string][]*multipart.FileHeader,
) (map[string]string, error) {
	result := map[string]string{}
	errChan := make(chan error, len(files))
	doneChan := make(chan struct{}, len(files))

	for fieldName, fileHeaders := range files {
		go func(fieldName string, fileHeaders []*multipart.FileHeader) {
			for _, fileHeader := range fileHeaders {
				s3URL, err := u.Upload(ctx, folder, fileHeader)
				if err != nil {
					errChan <- err
					return
				}
				result[fieldName] = s3URL
			}
			doneChan <- struct{}{}
		}(fieldName, fileHeaders)
	}

	for i := 0; i < len(files); i++ {
		select {
		case err := <-errChan:
			return nil, err
		case <-doneChan:
		}
	}

	return result, nil
}
