package awsutil

import (
	"context"
	"fmt"
	"src/common/setting"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func S3Client() *s3.Client {
	region := setting.S3_REGION
	cfg, _ := config.LoadDefaultConfig(
		context.TODO(),
		config.WithCredentialsProvider(
			credentials.NewStaticCredentialsProvider(
				setting.S3_ACCESS_KEY_ID, setting.S3_SECRET_ACCESS_KEY, "",
			)),
		config.WithRegion(region),
	)
	return s3.NewFromConfig(cfg, func(o *s3.Options) {
		o.BaseEndpoint = aws.String(fmt.Sprintf("https://%s.r2.cloudflarestorage.com", setting.S3_ACCOUNT_ID))
	})
}
