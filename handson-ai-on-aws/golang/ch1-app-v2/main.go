package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"strings"

	"ch1-app-v2/rekognition"
	"ch1-app-v2/storage"

	"github.com/aws/aws-sdk-go-v2/config"
)

func main() {
	bucketName := flag.String("b", "", "The bucket containing the image")

	flag.Parse()

	if *bucketName == "" {
		fmt.Println("You must supply a bucket (-b BUCKET)")
		return
	}

	objects, err := storage.GetAllFiles(bucketName)
	if err != nil {
		fmt.Print(err)
	}

	contents := objects.Contents
	for _, object := range contents {
		if strings.HasSuffix(*object.Key, ".jpg") {
			fmt.Printf("Objects detected in image %s:\n", *object.Key)
			labelsObject, err := rekognition.DetectObjects(bucketName, object.Key)
			if err != nil {
				fmt.Print(err)
			}

			labels := labelsObject.Labels
			for _, label := range labels {
				fmt.Printf("-- %s: %f\n", *label.Name, *label.Confidence)
			}
		}
	}

}

func init() {
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("ap-southeast-2"))
	if err != nil {
		log.Fatalf("failed to load configuration, %v", err)
	}

	// What would be a more idiomatic way to create these?
	storage.New(cfg)
	rekognition.New(cfg)
}
