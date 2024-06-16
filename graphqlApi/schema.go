package graphqlApi

import (
	"log"

	"github.com/graphql-go/graphql"
)


var	fields = graphql.Fields{
		"hello": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return "world", nil
			},
		},
	}
var rootQuery = graphql.ObjectConfig{Name: "RootQuery", Fields: fields}
var	schemaConfig = graphql.SchemaConfig{Query: graphql.NewObject(rootQuery)}

func CreateSchema() graphql.Schema {
	schema, err := graphql.NewSchema(schemaConfig)
	if err != nil {
		log.Fatalf("failed to create new schema, error: %v", err)
	}
	return schema
}