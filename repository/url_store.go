package repository

import (
	"context"
	"errors"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/expression"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/cockscomb/tinyurl/domain/entity"
	"time"
)

type URLStoreConfig struct {
	TableName string `env:"URL_TABLE,required"`
}

type URLStore struct {
	config *URLStoreConfig
	db     *dynamodb.Client
}

func NewURLStore(config *URLStoreConfig, db *dynamodb.Client) *URLStore {
	return &URLStore{config: config, db: db}
}

func (s *URLStore) Create(ctx context.Context, url *entity.TinyURL) error {
	row := URLRow{
		ID:        url.ID,
		URL:       url.URL.String(),
		CreatedAt: time.Now(),
	}
	av, err := attributevalue.MarshalMap(row)
	if err != nil {
		return err
	}
	expr, err := expression.NewBuilder().WithCondition(
		expression.AttributeNotExists(expression.Name("id")),
	).Build()
	if err != nil {
		return err
	}
	_, err = s.db.PutItem(ctx, &dynamodb.PutItemInput{
		TableName:                 aws.String(s.config.TableName),
		Item:                      av,
		ConditionExpression:       expr.Condition(),
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
	})
	if err != nil {
		var conditionalCheckFailed *types.ConditionalCheckFailedException
		if errors.As(err, &conditionalCheckFailed) {
			return entity.ErrAlreadyExists
		}
		return err
	}
	return nil
}

func (s *URLStore) Find(ctx context.Context, id string) (*entity.TinyURL, error) {
	output, err := s.db.GetItem(ctx, &dynamodb.GetItemInput{
		TableName: aws.String(s.config.TableName),
		Key: map[string]types.AttributeValue{
			"id": &types.AttributeValueMemberS{Value: id},
		},
	})
	if err != nil {
		return nil, err
	}
	if output.Item == nil {
		return nil, entity.ErrNotFound
	}
	var row URLRow
	if err := attributevalue.UnmarshalMap(output.Item, &row); err != nil {
		return nil, err
	}
	return row.ToEntity()
}
