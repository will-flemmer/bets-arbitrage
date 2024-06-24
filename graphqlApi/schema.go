package graphqlApi

import (
	"log"

	"github.com/graphql-go/graphql"
)


func CreateSchema() graphql.Schema {
	queryType := CreateQueryType()
	schemaConfig := graphql.SchemaConfig{Query: queryType}
	schema, err := graphql.NewSchema(schemaConfig)
	if err != nil {
		log.Fatalf("failed to create new schema, error: %v", err)
	}
	return schema
}