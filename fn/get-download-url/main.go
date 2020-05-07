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

const DefaultHTTPScheme = "https"

var cfDomain string
var downloadSvc *service.DownloadService

func init() {

	cfDomain = os.Getenv("CF_DOMAIN_NAME")

	downloadSvc = service.NewDownloadService()
}

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	reqData := &model.GetDownloadURLRequest{}

	err := apigwevent.ParseRequest(request, reqData)
	if err != nil {
		return apigwevent.BadRequest(err)
	}
	log.WithFields(log.Fields{"file": reqData.File}).Info("Get download url request")

	// these fields should not be open to frontend
	//
	reqData.Scheme = DefaultHTTPScheme
	reqData.Domain = cfDomain

	respData, err := downloadSvc.GetDownloadURL(reqData)
	if err != nil {
		return apigwevent.InternalServerError(err)
	}
	return apigwevent.OK(respData)
}

func main() {
	lambda.Start(handler)
}
