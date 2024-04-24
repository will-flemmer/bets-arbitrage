package main

import (
	"fmt"
	"scraping/bets"
	httpServer "scraping/http_server"
	"scraping/utils"
)

// Odds endpoint - https://the-odds-api.com/liveapi/guides/v4/#parameters-2
// /v4/sports/{sport}/odds/?apiKey={apiKey}&regions={regions}&markets={markets}&oddsFormat={oddsFormat}&commenceTimeFrom={commenceTimeFrom}&commenceTimeTo={commenceTimeTo}

func main() {
	// apiToken := env.API_TOKEN
	// bets.GetSports(apiToken)
	// getBets(apiToken)

	// wrangling.FetchAndStoreData()
	startHttpServer()
}

func startHttpServer() {
	go httpServer.StartHttpServer()
	var out string
	fmt.Scanln(&out)
}

func getBets(apiToken string) {
	sports := utils.LoadSports()
	for _, sport := range sports.Soccer {
		bets.FindBets(sport, apiToken)
	}
}