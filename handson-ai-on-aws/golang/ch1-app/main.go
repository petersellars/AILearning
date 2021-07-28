package main

import (
	"flag"
	"fmt"
	"strings"

	"ch1-app/rekognition"
	"ch1-app/storage"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
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
	creds := credentials.NewEnvCredentials()

	sess := session.New(&aws.Config{
		Credentials: creds,
		Region:      aws.String("ap-southeast-2"),
	})

	storage.New(sess)
	rekognition.New(sess)
}
