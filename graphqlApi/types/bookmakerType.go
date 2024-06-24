package types

import (
	"errors"
	"scraping/wrangling"

	"github.com/graphql-go/graphql"
)

var titleField = &graphql.Field{
	Type: graphql.String,
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		bookmaker, ok := p.Source.(wrangling.Bookmaker)
		if !ok {
			return bookmaker, errors.New("expected source to be a Bookmaker, but it was not")
		}
		return bookmaker.Title, nil
	},
}

var marketField = &graphql.Field{
	Type: graphql.NewList(MarketType),
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		bookmaker, ok := p.Source.(wrangling.Bookmaker)
		if !ok {
			return bookmaker, errors.New("expected source to be a Bookmaker, but it was not")
		}
		var markets []wrangling.Market
		wrangling.GlobalDB.Where("bookmaker_id = ?", bookmaker.ID).Find(&markets)

		return markets, nil
	},
}


var BookmakerType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "bookmaker",
		Description: "A bookmaker",
		Fields: graphql.Fields{
			"title": titleField,
			"markets": marketField,
		},
	},
)