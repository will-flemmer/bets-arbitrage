package bets

import (
	"fmt"
	"scraping/utils"
)


func GetSports(args ...string) {
	apiToken := args[0]
	endpoint := fmt.Sprintf("https://api.the-odds-api.com/v4/sports/?apiKey=%s", apiToken)

	var sports []utils.Sport
	// [utils.Sport] is a generic type
	utils.GetJson[utils.Sport](endpoint, &sports)
	var sportKeys []string

	for _, element := range sports {
		if element.Active == true {
			sportKeys = append(sportKeys, element.Key)
		}
	}

	// fmt.Println("total sports: ", len(sports))
	fmt.Println(sportKeys)
}
