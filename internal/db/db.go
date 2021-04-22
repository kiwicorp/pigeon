package db

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

var (
	sess   = session.Must(session.NewSession())
	client = dynamodb.New(sess)
)

// CreateItem creates a new item in the database.
func CreateItem(item TableItemAvMap) (*dynamodb.PutItemOutput, error) {
	av, err := item.AvMap()
	if err != nil {
		return nil, err
	}
	in := &dynamodb.PutItemInput{
		Item:      av,
		TableName: item.TableName(),
	}
	return client.PutItem(in)
}

// DeleteItem deletes an item from the database.
func DeleteItem(item TableItemKey) (*dynamodb.DeleteItemOutput, error) {
	return client.DeleteItem(&dynamodb.DeleteItemInput{
		TableName: item.TableName(),
		Key:       item.Key(),
	})
}

// GetItem fetches an item from the database.
func GetItem(item TableItemKey) (*dynamodb.GetItemOutput, error) {
	return client.GetItem(&dynamodb.GetItemInput{
		TableName: item.TableName(),
		Key:       item.Key(),
	})
}

// UpdateItem updates an item in the database.
func UpdateItem(item TableItemUpdate) (*dynamodb.UpdateItemOutput, error) {
	return client.UpdateItem(&dynamodb.UpdateItemInput{
		ExpressionAttributeValues: item.ExpressionAv(),
		TableName:                 item.TableName(),
		Key:                       item.Key(),
		ReturnValues:              aws.String(item.ReturnValues()),
		UpdateExpression:          aws.String(item.UpdateExpression()),
	})
}
