package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/rekognition"
	"github.com/aws/aws-sdk-go-v2/service/rekognition/types"
)

func main() {
	bucketName := flag.String("b", "", "The bucket containing the image")

	flag.Parse()

	if *bucketName == "" {
		fmt.Println("You must supply a bucket (-b BUCKET)")
		return
	}

	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("ap-southeast-2"))
	if err != nil {
		log.Fatalf("failed to load configuration, %v", err)
	}

	svc := rekognition.NewFromConfig(cfg)

	inputImage := "beagle.jpg"

	input := &rekognition.DetectLabelsInput{
		Image: &types.Image{
			S3Object: &types.S3Object{
				Bucket: bucketName,
				Name:   &inputImage,
			},
		},
	}

	output, err := svc.DetectLabels(context.TODO(), input)
	if err != nil {
		fmt.Print(err)
	}

	labels := output.Labels
	for _, label := range labels {
		fmt.Printf("-- %s: %f\n", *label.Name, *label.Confidence)
	}
}
