package wrangling

import (
	"fmt"
	"scraping/utils"
)

const (
	REGIONS = "uk"
	MARKETS = "h2h"
)

func FetchAndStoreData() error {
	fixtures, err := fetchFixtures()
	if err != nil {
		return err
	}

	dbHandle := LoadDb()
	for _, fixture := range fixtures {
		res := dbHandle.Create(&fixture)
		fmt.Printf("%+v", res.Error)
	}

	return err
}

func fetchFixtures() ([]Fixture, error) {
	apiToken := utils.LoadEnv().API_TOKEN
	println(apiToken)
	sports := utils.LoadSports()
	
	var fixtures []Fixture
	
	var err error	
	for _, sportKey := range sports.Soccer {
		endpoint := fmt.Sprintf(
			"https://api.the-odds-api.com/v4/sports/%s/odds/?apiKey=%s&regions=%s&markets=%s",
			sportKey,
			apiToken,
			REGIONS,
			MARKETS,
		)
		
		var f []Fixture
		err = utils.GetJson(endpoint, &f)
		fixtures = append(fixtures, f...)
	}

	return fixtures, err
}
