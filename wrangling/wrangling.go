package wrangling

// find best odds for each outcome accross all bookmakers
// will retun a slice which has 3 BestOdds structs, one for each outcome
func GetBestOdds(fixture Fixture, marketKey string) []BestOdds {
	var bestOdds []BestOdds
	for _, bookmaker := range fixture.Bookmakers {
		filteredMarkets := filterMarketsByKey(bookmaker.Markets, marketKey)
		for _, market := range filteredMarkets {
			handleMarket(&bestOdds, market.Outcomes, bookmaker.Title)
		}
	}
	return bestOdds
}

func handleMarket(bestOddsPointer *[]BestOdds, outcomes []Outcome, bookmakerName string) {
	if len(*bestOddsPointer) == 0 {
		for _, outcome := range outcomes {
			*bestOddsPointer = append(*bestOddsPointer, bestOddsStructFromOutcome(outcome, bookmakerName))
		}
		return
	}
	for _, outcome := range outcomes {
		checkIfBetterOdds(bestOddsPointer, outcome, bookmakerName)
	}
}

func checkIfBetterOdds(bestOddsPointer *[]BestOdds, outcome Outcome, bookmakerName string) {
	for index, odds := range *bestOddsPointer {
		if outcome.Name != odds.TeamName {
			continue
		}
		if outcome.Price > odds.Odds {
			(*bestOddsPointer)[index] = bestOddsStructFromOutcome(outcome, bookmakerName)
		}
	}
}

func bestOddsStructFromOutcome(outcome Outcome, bookmakerName string) BestOdds {
	return BestOdds{
		TeamName:      outcome.Name,
		Odds:          outcome.Price,
		BookmakerName: bookmakerName,
	}
}

func filterMarketsByKey(markets []Market, marketKey string) []Market {
	var filteredMarkets []Market
	for _, market := range markets {
		if market.Key == marketKey {
			filteredMarkets = append(filteredMarkets, market)
		}
	}
	return filteredMarkets
}
