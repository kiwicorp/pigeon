package content

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

var (
	sess           = session.Must(session.NewSession())
	dynamoDbClient = dynamodb.New(sess)
)

func putItem(item map[string]*dynamodb.AttributeValue, tableName string) (*dynamodb.PutItemOutput, error) {
	in := &dynamodb.PutItemInput{
		Item:      item,
		TableName: &tableName,
	}
	return dynamoDbClient.PutItem(in)
}
