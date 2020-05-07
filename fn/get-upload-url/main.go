package main

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	log "github.com/sirupsen/logrus"
	"os"
	"serverless-demo/awsapi/apigwevent"
	"serverless-demo/model"
	"serverless-demo/service"
)

var bucket string
var uploadSvc *service.UploadService

func init() {

	bucket = os.Getenv("S3_BUCKET")

	uploadSvc = service.NewUploadService()
}

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	reqData := &model.GetUploadURLRequest{}

	err := apigwevent.ParseRequest(request, reqData)
	if err != nil {
		return apigwevent.BadRequest(err)
	}
	log.WithFields(log.Fields{"file": reqData.File, "contentType": reqData.ContentType,
		"width": reqData.Width, "height": reqData.Height}).Info("Get upload url request")

	// these fields should not be open to frontend
	//
	reqData.Bucket = bucket

	respData, err := uploadSvc.GetUploadURL(reqData)
	if err != nil {
		return apigwevent.InternalServerError(err)
	}
	return apigwevent.OK(respData)
}

func main() {

	lambda.Start(handler)
}
