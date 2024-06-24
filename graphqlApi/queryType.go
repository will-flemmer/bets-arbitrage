package graphqlApi

import (
	"scraping/graphqlApi/types"
	"scraping/wrangling"

	"github.com/graphql-go/graphql"
)

func buildAllFixtures() *graphql.Field {
	fixturesArgs := graphql.FieldConfigArgument{
		"limit": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.Int),
		},
	}

	return &graphql.Field{
		Type: graphql.NewList(types.FixtureType),
		Args: fixturesArgs,
		Resolve: func(params graphql.ResolveParams) (interface{}, error) {
			limit, _ := params.Args["limit"].(int)

			var fixtures []wrangling.Fixture
			wrangling.GlobalDB.Limit(limit).Preload("Bookmakers").Find(&fixtures)
			return fixtures, nil
		},
	}
}

func CreateQueryType() *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"allfixtures": buildAllFixtures(),
		},
	})
}


