package rekognition

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/rekognition"
	"github.com/aws/aws-sdk-go-v2/service/rekognition/types"
)

var (
	rekognitionClient *rekognition.Client
)

func New(awsConfig aws.Config) {
	rekognitionClient = rekognition.NewFromConfig(awsConfig)
}

func DetectObjects(bucketName *string, imageName *string) (*rekognition.DetectLabelsOutput, error) {

	input := &rekognition.DetectLabelsInput{
		Image: &types.Image{
			S3Object: &types.S3Object{
				Bucket: bucketName,
				Name:   imageName,
			},
		},
	}

	return rekognitionClient.DetectLabels(context.TODO(), input)
}
