package db

import "github.com/aws/aws-sdk-go/service/dynamodb"

type TableItem interface {
	TableName() *string
}

type TableItemAvMap interface {
	TableItem

	AvMap() (map[string]*dynamodb.AttributeValue, error)
}

type TableItemKey interface {
	TableItem

	Key() map[string]*dynamodb.AttributeValue
}

type TableItemUpdate interface {
	TableItemKey

	ExpressionAv() map[string]*dynamodb.AttributeValue
	ReturnValues() string
	UpdateExpression() string
}
