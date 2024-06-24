package main

import (
	"scraping/config"
	httpServer "scraping/http_server"
	"scraping/wrangling"
)

// Odds endpoint - https://the-odds-api.com/liveapi/guides/v4/#parameters-2
// /v4/sports/{sport}/odds/?apiKey={apiKey}&regions={regions}&markets={markets}&oddsFormat={oddsFormat}&commenceTimeFrom={commenceTimeFrom}&commenceTimeTo={commenceTimeTo}

func main() {
	// apiToken := env.API_TOKEN
	// bets.GetSports(apiToken)
	// getBets(apiToken)
	// bets.FindBets()

	config.Init()
	wrangling.LoadDb()
	startHttpServer()
}

func startHttpServer() {
	httpServer.StartHttpServer()
}
