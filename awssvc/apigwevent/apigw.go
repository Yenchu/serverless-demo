package apigwevent

import (
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"log"
	"net/http"
)

func ParseRequest(req events.APIGatewayProxyRequest, data interface{}) error {

	if req.Body != "" {
		err := json.Unmarshal([]byte(req.Body), &data)
		if err != nil {
			log.Printf("json error: %v", err)
			return err
		}
	}
	return nil
}

func Response(body interface{}, status int) (events.APIGatewayProxyResponse, error) {

	bodyBytes, err := json.Marshal(body)
	if err != nil {
		log.Printf("json error: %v", err)
		return events.APIGatewayProxyResponse{}, err
	}

	bodyStr := string(bodyBytes)

	return events.APIGatewayProxyResponse{
		StatusCode: status,
		Body:       bodyStr,
	}, nil
}

func OK(body interface{}) (events.APIGatewayProxyResponse, error) {

	return Response(body, http.StatusOK)
}

func BadRequest(err error) (events.APIGatewayProxyResponse, error) {

	return Response(err, http.StatusBadRequest)
}

func InternalServerError(err error) (events.APIGatewayProxyResponse, error) {

	return Response(err, http.StatusInternalServerError)
}