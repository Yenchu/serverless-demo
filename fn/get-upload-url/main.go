package main

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"os"
	"serverless-demo/awssvc/apigwevent"
	"serverless-demo/model"
	"serverless-demo/service"
	"strconv"
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

	metadata := map[string]string{}

	width := reqData.Width
	if width > 0 {
		metadata["width"] = strconv.FormatInt(width, 10)
	}

	height := reqData.Height
	if height > 0 {
		metadata["height"] = strconv.FormatInt(height, 10)
	}

	respData, err := uploadSvc.GetUploadURL(bucket, reqData.File, reqData.ContentType, metadata)
	if err != nil {
		return apigwevent.InternalServerError(err)
	}
	return apigwevent.OK(respData)
}

func main() {

	lambda.Start(handler)
}
