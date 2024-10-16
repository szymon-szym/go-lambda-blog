package awssdkservice

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
)

type SecretsManagerService struct {
	client *secretsmanager.Client
}

func InitializeSecretsManager(cfg aws.Config) SecretsManagerService {

	client := secretsmanager.NewFromConfig(cfg)
	return SecretsManagerService{
		client: client,
	}
}

func (s SecretsManagerService) GetSecret(secretName string) (string, error) {
	input := &secretsmanager.GetSecretValueInput{
		SecretId:     aws.String(secretName),
		VersionStage: aws.String("AWSCURRENT"),
	}

	result, err := s.client.GetSecretValue(context.TODO(), input)
	if err != nil {
		return "", err
	}

	return *result.SecretString, nil
}
