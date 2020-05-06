package awssvc

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/external"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/dynamodbattribute"
)

func NewDynamoDBClient() *DynamoDBClient {

	awsCfg, err := external.LoadDefaultAWSConfig()
	if err != nil {
		panic("failed to load config, " + err.Error())
	}

	client := dynamodb.New(awsCfg)

	return &DynamoDBClient{
		client: client,
	}
}

type DynamoDBClient struct {
	client *dynamodb.Client
}

func (api *DynamoDBClient) GetItem(tableName, keyName, keyVal string, item interface{}) error {

	key := map[string]dynamodb.AttributeValue{
		keyName: {
			S: aws.String(keyVal),
		},
	}

	input := &dynamodb.GetItemInput{
		Key:       key,
		TableName: aws.String(tableName),
	}

	req := api.client.GetItemRequest(input)

	resp, err := req.Send(context.Background())
	if err != nil {
		return err
	}
	return dynamodbattribute.UnmarshalMap(resp.Item, item)
}

func (api *DynamoDBClient) PutItem(tableName string, itemVal interface{}) error {

	item, err := dynamodbattribute.MarshalMap(itemVal)
	if err != nil {
		return err
	}

	input := &dynamodb.PutItemInput{
		Item:      item,
		TableName: aws.String(tableName),
	}

	req := api.client.PutItemRequest(input)

	_, err = req.Send(context.Background())
	return err
}
