package storage

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

var (
	s3Client *s3.Client
)

func New(awsConfig aws.Config) {
	s3Client = s3.NewFromConfig(awsConfig)
}

func GetAllFiles(bucketName *string) (*s3.ListObjectsV2Output, error) {
	listObjectsInput := &s3.ListObjectsV2Input{
		Bucket: bucketName,
	}
	return s3Client.ListObjectsV2(context.TODO(), listObjectsInput)
}
