package graphqlApi

import (
	"log"

	"github.com/graphql-go/graphql"
	"gorm.io/gorm"
)


func CreateSchema(db *gorm.DB) graphql.Schema {
	queryType := CreateQueryType(db)
	schemaConfig := graphql.SchemaConfig{Query: queryType}
	schema, err := graphql.NewSchema(schemaConfig)
	if err != nil {
		log.Fatalf("failed to create new schema, error: %v", err)
	}
	return schema
}