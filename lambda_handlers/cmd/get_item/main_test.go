package main

import (
	"context"
	"errors"
	api "lambda_handlers/api"
	"testing"

	"github.com/aws/aws-lambda-go/events"
	"github.com/stretchr/testify/assert"
)

type MockGetItemsService struct {
	mockGetItem func(id int) (*api.Item, error)
}

func (m *MockGetItemsService) GetItem(id int) (*api.Item, error) {
	return m.mockGetItem(id)
}

func TestLambdaHandler_Success(t *testing.T) {
	mockService := &MockGetItemsService{
		mockGetItem: func(id int) (*api.Item, error) {
			return &api.Item{Id: id, Name: "Test Item", Price: 42.11}, nil
		},
	}

	handler := InitializeLambdaHandler(mockService)

	req := events.APIGatewayProxyRequest{
		PathParameters: map[string]string{
			"id": "1",
		},
	}

	resp, err := handler.HandleRequest(context.TODO(), req)

	assert.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, "{\"id\":1,\"name\":\"Test Item\",\"price\":42.11}", resp.Body)
}

func TestLambdaHandler_NotFound(t *testing.T) {
	mockService := &MockGetItemsService{
		mockGetItem: func(id int) (*api.Item, error) {
			return nil, errors.New("Item not found")
		},
	}

	handler := InitializeLambdaHandler(mockService)

	req := events.APIGatewayProxyRequest{
		PathParameters: map[string]string{
			"id": "1",
		},
	}

	resp, err := handler.HandleRequest(context.TODO(), req)

	assert.NoError(t, err)
	assert.Equal(t, 404, resp.StatusCode)
	assert.Equal(t, "Item not found", resp.Body)
}

func TestLambdaHandler_InvalidId(t *testing.T) {
	mockService := &MockGetItemsService{
		mockGetItem: func(id int) (*api.Item, error) {
			return &api.Item{Id: id, Name: "Test Item", Price: 42.11}, nil
		},
	}

	handler := InitializeLambdaHandler(mockService)

	req := events.APIGatewayProxyRequest{
		PathParameters: map[string]string{
			"id": "invalid",
		},
	}

	resp, err := handler.HandleRequest(context.TODO(), req)

	assert.NoError(t, err)
	assert.Equal(t, 400, resp.StatusCode)
	assert.Equal(t, "id must be an integer", resp.Body)
}
