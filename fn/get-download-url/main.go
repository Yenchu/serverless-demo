package main

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"os"
	"serverless-demo/awssvc/apigwevent"
	"serverless-demo/model"
	"serverless-demo/service"
)

var domainName string
var downloadSvc *service.DownloadService

func init() {

	domainName = os.Getenv("CF_DOMAIN_NAME")

	downloadSvc = service.NewDownloadService()
}

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	reqData := &model.GetDownloadURLRequest{}

	err := apigwevent.ParseRequest(request, reqData)
	if err != nil {
		return apigwevent.BadRequest(err)
	}

	file := reqData.File

	respData, err := downloadSvc.GetDownloadURL(domainName, file)
	if err != nil {
		return apigwevent.InternalServerError(err)
	}
	return apigwevent.OK(respData)
}

func main() {
	lambda.Start(handler)
}
