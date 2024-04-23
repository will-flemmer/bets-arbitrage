package bets

import (
	"errors"
	"fmt"
	"scraping/utils"
	"scraping/wrangling"
)

// example cases
// cash: $100
// BookmakersA:
// - SA, price 2,
// - Tonga, price 30,
// - Draw, price 100

// BookmakersB:
// - SA, price 1,
// - Tonga, price 50,
// - Draw, price 100

// Best SA bet is 2
// Best tonga bet is 50
// Best Draw bet is  100
// totalOdds = 2 + 50 + 100 = 152
// 152 / 2 = 76 odd units
// 152 / 50 = 3.04 odd units
// 152 / 100 = 1.52 odd units

// totalBets = 76 + 3.04 + 1.52 = 80.56
// $/bet unit = $100/80.56 = $1.24/unit

// !!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!
// POTENTIAL_WINNINGS = diff(totalOdds, totalBets)
// !!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!

// SA bet = 76 units at $1.24/unit = $94.24
// for a return of $94.24 * 2 = $188.48

// Tonga bet = 3.04 units at $1.24/unit = $3.76
// for a return of $3.76 * 50 = $188

// Draw bet = 1.52 units at $1.24/unit = $1.88
// for a return of $1.88 * 100 = $188

// This is a scenario where I should place these bets, i.e minReturn > cash * 1.2

// Find how much to bet on each outcome to maximise profit
//

func FindBets(sportKey string, apiToken string) {
	var availableCash float32 = 100
	regions := "uk"
	markets := "h2h"
	endpoint := fmt.Sprintf("https://api.the-odds-api.com/v4/sports/%s/odds/?apiKey=%s&regions=%s&markets=%s", sportKey, apiToken, regions, markets)
	var fixtures []wrangling.Fixture

	println("Profitable bets will be printed below:")
	utils.GetJson(endpoint, &fixtures)
	for _, fixture := range fixtures {
		if len(fixture.Bookmakers) < 2 {
			continue
		}
		bestOddsForGame := wrangling.GetBestOdds(fixture, markets)
		bets, err := calculateBets(&bestOddsForGame, availableCash)

		if err != nil {
			continue
		}
		prettyPrintProfitableBet(&bets, availableCash)
	}
}

func prettyPrintProfitableBet(bets *[]Bet, cashInput float32) {
	avgReturn := calcAvgReturn(bets)
	percentageReturn := 100 * (avgReturn - cashInput) / cashInput

	fmt.Println("Bet is Profitable!")
	fmt.Println("With a cash input of", cashInput, "you will get a return of", avgReturn, "----", percentageReturn, "%")
}

func calcAvgReturn(bets *[]Bet) float32 {
	var sum float32 = 0
	for _, bet := range *bets {
		sum = sum + bet.cashReturn
	}
	return float32(sum) / float32(len(*bets))
}

type Bet struct {
	team           string
	odds           float32
	oddsUnits 		 float32
	cashInvestment float32 // this needs to be rounded to nearest whole number
	bookmaker      string
	cashReturn		 float32
}

// Need to round to nearest whole number when calculating cash to invest
func calculateBets(outcomesPointer *[]wrangling.BestOdds, cashAvailable float32) ([]Bet, error) {
	totalOdds := calcTotalOdds(outcomesPointer)
	normalizedTotalOdds := normalizedTotalOdds(totalOdds, outcomesPointer)
	bets := calcCashSplit(totalOdds, normalizedTotalOdds, outcomesPointer, cashAvailable)
	if isProfitable(bets, cashAvailable) {
		return bets, nil
	}
	return nil, errors.New("fixture is not profitable")
}

func calcCashSplit(totalOdds float32, normalizedTotalOdds float32, outcomesPointer *[]wrangling.BestOdds, cashAvailable float32) []Bet {
	var bets []Bet
	cashPerOdd := cashAvailable / normalizedTotalOdds
	for _, outcome := range *outcomesPointer {
		oddUnits := totalOdds / outcome.Odds

		cashInvestment := oddUnits * cashPerOdd
		cashReturn := cashInvestment * outcome.Odds
		bet := Bet{
			team: outcome.TeamName,
			odds: outcome.Odds,
			cashInvestment: cashInvestment, // this needs to be rounded to nearest whole number
			bookmaker: outcome.BookmakerName,
			cashReturn: cashReturn }
		bets = append(bets, bet)
	}
	return bets
}

func oddsToCash(totalOdds float32, cashAvailable float32, outcomeOdds float32) float32 {
	cashPerOddPoint := cashAvailable / totalOdds
	return outcomeOdds * cashPerOddPoint
}

func normalizedTotalOdds(totalOdds float32, outcomesPointer *[]wrangling.BestOdds) float32 {
	var sum float32 = 0
	for _, outcome := range *outcomesPointer {
		sum = sum + (totalOdds / outcome.Odds)
	}
	return sum
}

func isProfitable(allBets []Bet, cashInput float32) bool {
	for _, bet := range allBets {
		if bet.cashReturn <= cashInput {
			return false
		}
	}
	return true
}

func calcTotalOdds(oddsSlice *[]wrangling.BestOdds) float32 {
	var sum float32
	for _, outcome := range *oddsSlice {
		sum = sum + outcome.Odds
	}
	return sum
}
