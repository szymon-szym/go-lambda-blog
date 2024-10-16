package main

import (
	"context"
	"encoding/json"
	api "lambda_handlers/api"
	awssdkconfig "lambda_handlers/internal/awssdk"
	"lambda_handlers/internal/db"
	"lambda_handlers/internal/helpers"
	"lambda_handlers/internal/items"
	"log"
	"os"
	"strconv"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type GetItemsService interface {
	GetItem(id int) (*api.Item, error)
}

type LambdaHandler struct {
	svc GetItemsService
}

func InitializeLambdaHandler(svc GetItemsService) *LambdaHandler {
	return &LambdaHandler{
		svc: svc,
	}
}

func (h *LambdaHandler) HandleRequest(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	strId, err := helpers.GetPathParameter("id", request.PathParameters)

	if err != nil {
		return helpers.SendResponse("id is required", 400), nil
	}

	id, err := strconv.Atoi(*strId)

	if err != nil {
		return helpers.SendResponse("id must be an integer", 400), nil
	}

	log.Printf("id: %d", id)

	result, err := h.svc.GetItem(id)

	if err != nil {
		if err.Error() == "Item not found" {
			return helpers.SendResponse("Item not found", 404), nil
		}
		return helpers.SendResponse("error", 500), nil
	}

	jsonRes, err := json.Marshal(result)

	if err != nil {
		log.Printf("error marshalling response: %s", err.Error())
		return helpers.SendResponse("internal server error", 500), nil
	}

	return helpers.SendResponse(string(jsonRes), 200), nil
}

func main() {

	dbSecretName := os.Getenv("DB_SECRET_NAME")

	log.Printf("dbSecretName: %s", dbSecretName)

	cfg, err := awssdkconfig.InitializeSdkConfig()

	if err != nil {
		log.Fatal(err)
	}

	secretsClient := awssdkconfig.InitializeSecretsManager(cfg)

	connString, err := secretsClient.GetSecret(dbSecretName)

	if err != nil {
		log.Fatal(err)
	}

	conn, err := db.InitializeDB(connString)

	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()

	log.Println("successfully connected to db")

	svc := items.InitializeItemsService(conn)

	handler := InitializeLambdaHandler(svc)

	lambda.Start(handler.HandleRequest)
}
