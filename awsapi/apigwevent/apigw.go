package apigwevent

import (
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func ParseRequest(req events.APIGatewayProxyRequest, data interface{}) error {

	if req.Body != "" {
		err := json.Unmarshal([]byte(req.Body), &data)
		if err != nil {
			return err
		}
	}
	return nil
}

func Response(body interface{}, status int) (events.APIGatewayProxyResponse, error) {

	bodyBytes, err := json.Marshal(body)
	if err != nil {
		log.WithFields(log.Fields{"error": err}).Error("Parse JSON failed")
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Body: err.Error(),
		}, err
	}

	bodyStr := string(bodyBytes)

	if status < 200 || status >= 400 {
		log.WithFields(log.Fields{"status": status, "body": bodyStr}).Error("REST error response")
	}

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
