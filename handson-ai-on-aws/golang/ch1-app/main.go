package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/rekognition"
)

func main() {

	// AWS Rekognition Detect Labels from AWS SDK for Go
	creds := credentials.NewEnvCredentials()

	sess := session.New(&aws.Config{
		Credentials: creds,
		Region:      aws.String("ap-southeast-2"),
	})

	svc := rekognition.New(sess)

	bucket := "contents.catosplace.ai-39c2f3a"
	inputImage := "beagle.jpg"

	input := &rekognition.DetectLabelsInput{
		Image: &rekognition.Image{
			S3Object: &rekognition.S3Object{
				Bucket: &bucket,
				Name:   &inputImage,
			},
		},
	}

	output, err := svc.DetectLabels(input)
	if err != nil {
		fmt.Print(err)
	}

	fmt.Print(output)
}
