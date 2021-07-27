package main

import (
	"flag"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/rekognition"
)

func main() {
	bucketName := flag.String("b", "", "The bucket containing the image")

	flag.Parse()

	if *bucketName == "" {
		fmt.Println("You must supply a bucket (-b BUCKET)")
		return
	}

	// AWS Rekognition Detect Labels from AWS SDK for Go
	creds := credentials.NewEnvCredentials()

	sess := session.New(&aws.Config{
		Credentials: creds,
		Region:      aws.String("ap-southeast-2"),
	})

	svc := rekognition.New(sess)

	inputImage := "beagle.jpg"

	input := &rekognition.DetectLabelsInput{
		Image: &rekognition.Image{
			S3Object: &rekognition.S3Object{
				Bucket: bucketName,
				Name:   &inputImage,
			},
		},
	}

	output, err := svc.DetectLabels(input)
	if err != nil {
		fmt.Print(err)
	}

	labels := output.Labels
	for _, label := range labels {
		fmt.Printf("-- %s: %f\n", *label.Name, *label.Confidence)
	}

}
