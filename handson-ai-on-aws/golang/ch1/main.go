package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/s3"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

const (
	CONTENTS_S3BUCKET_NAME = "contents.catosplace.ai"
	DATA_S3BUCKET_NAME     = "data.catosplace.ai"
	WEBSITE_S3BUCKET_NAME  = "website.catosplace.ai"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {

		// S3 Bucket Args
		// Block Public Access supported by PublicAccessBlock
		s3BucketArgs := &s3.BucketArgs{
			Acl: pulumi.String("private"),
		}

		// Create the contents, data and website buckets
		contentsBucket, _ := createS3Bucket(ctx, CONTENTS_S3BUCKET_NAME, s3BucketArgs, "Contents")
		dataBucket, _ := createS3Bucket(ctx, DATA_S3BUCKET_NAME, s3BucketArgs, "Data")
		websiteBucket, _ := createS3Bucket(ctx, WEBSITE_S3BUCKET_NAME, s3BucketArgs, "Website")

		// Export the name of the bucket
		exportS3BucketName(ctx, contentsBucket, "Contents")
		exportS3BucketName(ctx, dataBucket, "Data")
		exportS3BucketName(ctx, websiteBucket, "Website")

		// Upload Image to S3 Contents Bucket
		_, err := s3.NewBucketObject(ctx, "ContentsBucketBeagleObject", &s3.BucketObjectArgs{
			Key:    pulumi.String("beagle.jpg"),
			Bucket: contentsBucket.ID(),
			Source: pulumi.NewFileAsset("assets/beagle.jpg"),
		})
		if err != nil {
			return err
		}

		return nil
	})
}

func createS3Bucket(ctx *pulumi.Context, name string, s3BucketArgs *s3.BucketArgs, exportPrefix string) (*s3.Bucket, error) {
	bucket, err := s3.NewBucket(ctx, name, s3BucketArgs)
	if err != nil {
		return nil, err
	}
	_, err = s3.NewBucketPublicAccessBlock(ctx, exportPrefix+"BucketPublicAccessBlock", &s3.BucketPublicAccessBlockArgs{
		Bucket:                bucket.ID(),
		BlockPublicAcls:       pulumi.Bool(true),
		BlockPublicPolicy:     pulumi.Bool(true),
		IgnorePublicAcls:      pulumi.Bool(true),
		RestrictPublicBuckets: pulumi.Bool(true),
	})
	if err != nil {
		return nil, err
	}
	return bucket, nil
}

func exportS3BucketName(ctx *pulumi.Context, bucket *s3.Bucket, exportPrefix string) {
	ctx.Export(exportPrefix+"BucketName", bucket.ID())
}
