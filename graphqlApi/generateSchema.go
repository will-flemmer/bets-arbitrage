package graphqlApi

import (
	"log"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/graphql/testutil"
	"gorm.io/gorm"
)

func GenerateSchema(db *gorm.DB) interface{} {
	log.Println("Generating schema")
	result := graphql.Do(graphql.Params{
			Schema:        CreateSchema(db),
			RequestString: testutil.IntrospectionQuery,
	})
	return result.Data
}