package helpers

import (
	"fmt"

	"github.com/aws/aws-lambda-go/events"
)

func SendResponse(text string, code int) events.APIGatewayProxyResponse {
	return events.APIGatewayProxyResponse{
		StatusCode: code,
		Body:       text,
	}
}

func GetPathParameter(name string, params map[string]string) (*string, error) {
	if param, ok := params[name]; ok {
		return &param, nil
	} else {
		return nil, fmt.Errorf("missing path parameter: %s", name)
	}
}
