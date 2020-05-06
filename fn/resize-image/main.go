package main

import (
	"context"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"serverless-demo/service"
)

var imageSvc *service.ImageService

func init() {

	imageSvc = service.NewImageService()
}

func handler(ctx context.Context, s3Event events.S3Event) {

	for _, record := range s3Event.Records {

		s3 := record.S3

		bucket := s3.Bucket.Name
		key := s3.Object.Key

		fmt.Printf("[%s: %s - %s] Bucket = %s, Key = %s \n",
			record.EventSource, record.EventName, record.EventTime, bucket, key)

		//if record.EventName != "ObjectCreated:Put" {
		//	continue
		//}

		err := imageSvc.ResizeImage(bucket, key)
		if err != nil {
			fmt.Printf("ResizeImage failed: %v", err)
		}
	}
}

func main() {

	lambda.Start(handler)
}
