package rekognition

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/rekognition"
)

var (
	rekognitionClient *rekognition.Rekognition
)

func New(sess *session.Session) {
	rekognitionClient = rekognition.New(sess)
}

func DetectObjects(bucketName *string, imageName *string) (*rekognition.DetectLabelsOutput, error) {

	input := &rekognition.DetectLabelsInput{
		Image: &rekognition.Image{
			S3Object: &rekognition.S3Object{
				Bucket: bucketName,
				Name:   imageName,
			},
		},
	}

	return rekognitionClient.DetectLabels(input)
}
