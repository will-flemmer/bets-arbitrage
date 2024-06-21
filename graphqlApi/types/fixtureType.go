package types

import (
	"errors"
	"scraping/wrangling"

	"github.com/graphql-go/graphql"
)

var awayTeamField = &graphql.Field{
	Type: graphql.String,
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		fixture, ok := p.Source.(wrangling.Fixture)
		if !ok {
			return fixture, errors.New("expected source to be a Fixture, but it was not")
		}
		return fixture.AwayTeam, nil
	},
}


var bookmakersField = &graphql.Field{
	Type: graphql.NewList(BookmakerType),
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		fixture, ok := p.Source.(wrangling.Fixture)
		if !ok {
			return fixture, errors.New("expected source to be a Fixture, but it was not")
		}
		return fixture.Bookmakers, nil
	},
}

var FixtureType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "fixture",
		Description: "A sporting fixture",
		Fields: graphql.Fields{
			"awayTeam": awayTeamField,
			"bookmakers": bookmakersField,
		},
	},
)
