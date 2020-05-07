package main

import (
	"context"
	"fmt"
	log "github.com/sirupsen/logrus"
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

		log.WithFields(log.Fields{"source": record.EventSource, "event": record.EventName,
			"bucket": bucket, "key": key}).Info("S3 event record")

		//if record.EventName != "ObjectCreated:Put" {
		//	continue
		//}

		err := imageSvc.ResizeImage(bucket, key)
		if err != nil {
			log.WithFields(log.Fields{"error": err}).Error("ResizeImage failed")
		}
	}
}

func main() {

	lambda.Start(handler)
}
