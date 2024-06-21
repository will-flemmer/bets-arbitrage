package graphqlApi

import (
	"scraping/graphqlApi/types"
	"scraping/wrangling"

	"github.com/graphql-go/graphql"
	"gorm.io/gorm"
)

func buildAllFixtures(db *gorm.DB) *graphql.Field {
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
			db.Limit(limit).Find(&fixtures)
			return fixtures, nil
		},
	}
}

func CreateQueryType(db *gorm.DB) *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"allfixtures": buildAllFixtures(db),
		},
	})
}
