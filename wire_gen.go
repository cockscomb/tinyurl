// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"context"
	"github.com/cockscomb/tinyurl/repository"
	"github.com/cockscomb/tinyurl/usecase"
	"github.com/cockscomb/tinyurl/web"
	"github.com/cockscomb/tinyurl/web/controller"
)

// Injectors from wire.go:

func InitializeServer(ctx context.Context) (*web.Server, error) {
	config, err := ParseEnv()
	if err != nil {
		return nil, err
	}
	serverConfig := &config.ServerConfig
	templateConfig := &config.TemplateConfig
	template := web.NewTemplate(templateConfig)
	urlStoreConfig := &config.URLStoreConfig
	awsConfig := &config.AWSConfig
	config2, err := LoadAWSConfig(ctx, awsConfig)
	if err != nil {
		return nil, err
	}
	dynamoDBConfig := &config.DynamoDBConfig
	client := NewDynamoDBClient(config2, dynamoDBConfig)
	urlStore := repository.NewURLStore(urlStoreConfig, client)
	tinyURLUsecase := usecase.NewTinyURLUsecase(urlStore)
	controllerController := controller.NewController(tinyURLUsecase)
	server := web.NewServer(serverConfig, template, controllerController)
	return server, nil
}
