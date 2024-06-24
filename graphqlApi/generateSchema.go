package graphqlApi

import (
	"log"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/graphql/testutil"
)

func GenerateSchema() interface{} {
	log.Println("Generating schema")
	result := graphql.Do(graphql.Params{
			Schema:        CreateSchema(),
			RequestString: testutil.IntrospectionQuery,
	})
	return result.Data
}