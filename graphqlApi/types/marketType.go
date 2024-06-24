package types

import (
	"errors"
	"scraping/wrangling"

	"github.com/graphql-go/graphql"
)

var keyField = &graphql.Field{
	Type: graphql.String,
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		market, ok := p.Source.(wrangling.Market)
		if !ok {
			return market, errors.New("expected source to be a Market, but it was not")
		}
		return market.Key, nil
	},
}

var MarketType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "market",
		Description: "A Market",
		Fields: graphql.Fields{
			"key": keyField,
		},
	},
)