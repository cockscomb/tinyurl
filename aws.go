package main

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/google/wire"
)

type AWSConfig struct {
	Region             string `env:"REGION" envDefault:"ap-northeast-1"`
	UseDummyCredential bool   `env:"USE_DUMMY_CREDENTIAL" envDefault:"false"`
}

func LoadAWSConfig(ctx context.Context, cfg *AWSConfig) (aws.Config, error) {
	var opts []func(*config.LoadOptions) error
	opts = append(opts, config.WithRegion(cfg.Region))
	if cfg.UseDummyCredential {
		opts = append(opts, config.WithCredentialsProvider(
			credentials.NewStaticCredentialsProvider("dummy", "dummy", "dummy"),
		))
	}
	return config.LoadDefaultConfig(ctx, opts...)
}

type DynamoDBConfig struct {
	Endpoint string `env:"ENDPOINT"`
}

func NewDynamoDBClient(cfg aws.Config, dynamodbConfig *DynamoDBConfig) *dynamodb.Client {
	var opts []func(*dynamodb.Options)
	if dynamodbConfig.Endpoint != "" {
		opts = append(opts, dynamodb.WithEndpointResolver(
			dynamodb.EndpointResolverFromURL(dynamodbConfig.Endpoint)),
		)
	}
	return dynamodb.NewFromConfig(cfg, opts...)
}

var AWSSet = wire.NewSet(
	LoadAWSConfig,
	NewDynamoDBClient,
)
