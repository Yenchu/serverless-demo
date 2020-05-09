package main

import (
	"context"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	log "github.com/sirupsen/logrus"
	"os"
	"serverless-demo/awsapi/apigwevent"
	"serverless-demo/model"
	"serverless-demo/service"
)

var authSvc *service.AuthService

func init() {

	userPoolClientID := os.Getenv("USER_POOL_CLIENT_ID")

	cfg := &service.AuthServiceConfig{
		UserPoolClientID: userPoolClientID,
	}

	authSvc = service.NewAuthService(cfg)
}

func handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	reqData := &model.SignInRequest{}
	err := apigwevent.ParseRequest(req, reqData)
	if err != nil {
		return apigwevent.BadRequest(err)
	}
	log.WithFields(log.Fields{"username": reqData.Username}).Info("User signin")

	respData, err := authSvc.SignIn(reqData)
	if err != nil {
		return apigwevent.InternalServerError(err)
	}
	return apigwevent.OK(respData)
}

func main() {

	lambda.Start(handler)
}
