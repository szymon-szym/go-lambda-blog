package awssdkservice

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
)

func InitializeSdkConfig() (aws.Config, error) {
	return config.LoadDefaultConfig(context.TODO(), config.WithRegion("us-east-1"))
}
