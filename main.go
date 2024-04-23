package main

import (
	"fmt"
	"scraping/bets"
	"scraping/utils"
)

// Odds endpoint - https://the-odds-api.com/liveapi/guides/v4/#parameters-2
// /v4/sports/{sport}/odds/?apiKey={apiKey}&regions={regions}&markets={markets}&oddsFormat={oddsFormat}&commenceTimeFrom={commenceTimeFrom}&commenceTimeTo={commenceTimeTo}

func main() {
	env := utils.LoadEnv()
	fmt.Printf("%+v", env)
	apiToken := env.API_TOKEN
	// bets.GetSports(apiToken)
	getBets(apiToken)
}

func getBets(apiToken string) {
	sports := utils.LoadSports()
	for _, sport := range sports.Soccer {
		bets.FindBets(sport, apiToken)
	}
}