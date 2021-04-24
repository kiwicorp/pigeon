package db

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

var (
	sess   = session.Must(session.NewSession())
	client = dynamodb.New(sess)
)

// PutItem creates a new item in the database.
func PutItem(item TableItem) (*dynamodb.PutItemOutput, error) {
	av, err := dynamodbattribute.MarshalMap(item)
	if err != nil {
		return nil, err
	}
	in := &dynamodb.PutItemInput{
		Item:      av,
		TableName: item.TableName(),
	}
	return client.PutItem(in)
}

// GetItem fetches an item from the database.
func GetItem(item TableItemKey) error {
	result, err := client.GetItem(&dynamodb.GetItemInput{
		TableName: item.TableName(),
		Key:       item.Key(),
	})
	if err != nil {
		return err
	}
	err = dynamodbattribute.UnmarshalMap(result.Item, item)
	if err != nil {
		return err
	}
	return nil
}
