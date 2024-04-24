package wrangling

import (
	"fmt"
	"scraping/utils"
)

const (
	regions = "uk"
	markets = "h2h"
)

func FetchAndStoreData() error {
	fixtures, err := fetchFixtures()
	if err != nil {
		return err
	}

	firstFixtures := fixtures[:2]

	dbHandle := LoadDb()
	for _, fixture := range firstFixtures {
		res := dbHandle.Create(&fixture)
		fmt.Printf("%+v", res.Error)
	}

	return err
}

func fetchFixtures() ([]Fixture, error) {
	apiToken := utils.LoadEnv().API_TOKEN
	println(apiToken)
	sportKey := "soccer_austria_bundesliga"
	endpoint := fmt.Sprintf("https://api.the-odds-api.com/v4/sports/%s/odds/?apiKey=%s&regions=%s&markets=%s", sportKey, apiToken, regions, markets)

	var fixtures []Fixture
	err := utils.GetJson(endpoint, &fixtures)
	// println(err.Error())
	// var err error

	return fixtures, err
}


// https://api.the-odds-api.com/v4/sports/soccer_austria_bundesliga/odds/?apiKey=df39f2db5672e910241a1776437ca1f1&regions=uk&markets=h2h
// df39f2db5672e910241a1776437ca1f1