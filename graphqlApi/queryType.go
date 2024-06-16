package graphqlApi

import (
	"github.com/graphql-go/graphql"
)

var QueryType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Query",
	Fields: graphql.Fields{
			"latestPost": &graphql.Field{
					Type: graphql.String,
					Resolve: func(p graphql.ResolveParams) (interface{}, error) {
							return "Hello World!", nil
					},
			},
	},
})

var QueryTypeConfig = graphql.ObjectConfig{Name: "RootQuery", Fields: QueryType}