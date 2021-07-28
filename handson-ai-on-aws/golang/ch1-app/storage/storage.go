package storage

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

var (
	s3Client *s3.S3
)

func New(sess *session.Session) {
	s3Client = s3.New(sess)
}

func GetAllFiles(bucketName *string) (*s3.ListObjectsOutput, error) {
	listObjectsInput := &s3.ListObjectsInput{
		Bucket: bucketName,
	}
	return s3Client.ListObjects(listObjectsInput)
}
