package bets

import (
	"errors"
	"fmt"
	"scraping/wrangling"
	"time"
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

type ProfitableBet struct {
	FixtureName string `json:"fixtureName"`
	AvgReturnPercentage float32 `json:"avgReturnPercentage"`
	Outcomes []Bet `json:"outcomes"`
}

func FindBets() []ProfitableBet {
	var availableCash float32 = 100
	var fixtures []wrangling.Fixture

	println("Profitable bets will be printed below:")
	dbHandle := wrangling.LoadDb()
	currentDateTime := time.Now().Format(time.RFC1123)

	dbHandle.Where(
		"created_at < ?", currentDateTime,
	).Preload("Bookmakers.Markets.Outcomes").Find(&fixtures)

	// fmt.Printf("%+v", fixtures)
	var profitableBets []ProfitableBet
	for _, fixture := range fixtures {
		if len(fixture.Bookmakers) < 2 {
			continue
		}
		bestOddsForGame := wrangling.GetBestOdds(fixture, wrangling.MARKETS)
		// wrangling.GetBestOdds(fixture, wrangling.MARKETS)
		bets, err := calculateBets(&bestOddsForGame, availableCash)
		fmt.Printf("%+v\n", bets)
		
		if err != nil {
			println(err.Error())
			continue
		}
		profitableBets = append(profitableBets, createProfitableBet(&fixture, bets))
		
		// prettyPrintProfitableBet(&bets, availableCash)
	}
	
	fmt.Printf("\n%+v\n", profitableBets)
	return profitableBets
}

func createProfitableBet(fixture *wrangling.Fixture, bets []Bet) ProfitableBet {
	return ProfitableBet{
		FixtureName: fmt.Sprintf("Away Team = %s", fixture.AwayTeam),
		AvgReturnPercentage: calcAvgReturn(&bets),
		Outcomes: bets,
	}

}

// func AnalyseFixtures(fixtures *[]wrangling.Fixture, availableCash float32) ProfitableBet {
// 	for _, fixture := range *fixtures {
// 		if len(fixture.Bookmakers) < 2 {
// 			continue
// 		}
// 		bestOddsForGame := wrangling.GetBestOdds(fixture, wrangling.MARKETS)
// 		bets, err := calculateBets(&bestOddsForGame, availableCash)

// 		if err != nil {
// 			continue
// 		}
// 		prettyPrintProfitableBet(&bets, availableCash)
// 		profitableBets = append(profitableBets, bets)
// 	}
// 	return profitableBets
// }

func prettyPrintProfitableBet(bets *[]Bet, cashInput float32) {
	avgReturn := calcAvgReturn(bets)
	percentageReturn := 100 * (avgReturn - cashInput) / cashInput

	fmt.Println("Bet is Profitable!")
	fmt.Println("With a cash input of", cashInput, "you will get a return of", avgReturn, "----", percentageReturn, "%")
}

func calcAvgReturn(bets *[]Bet) float32 {
	var sum float32 = 0
	for _, bet := range *bets {
		sum = sum + bet.CashReturn
	}
	return float32(sum) / float32(len(*bets)) 
}

type Bet struct {
	Team           string `json:"team"`
	Odds           float32 `json:"odds"`
	CashInvestment float32 `json:"cashIn"` // this needs to be rounded to nearest whole number
	Bookmaker      string `json:"bookmaker"`
	CashReturn		 float32 `json:"cashReturn"`
}

// Need to round to nearest whole number when calculating cash to invest
func calculateBets(outcomesPointer *[]wrangling.BestOdds, cashAvailable float32) ([]Bet, error) {
	totalOdds := calcTotalOdds(outcomesPointer)
	normalizedTotalOdds := normalizedTotalOdds(totalOdds, outcomesPointer)
	bets := calcCashSplit(totalOdds, normalizedTotalOdds, outcomesPointer, cashAvailable)
	fmt.Printf("%+v", bets)
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
		CashReturn := cashInvestment * outcome.Odds
		bet := Bet{
			Team: outcome.TeamName,
			Odds: outcome.Odds,
			CashInvestment: cashInvestment, // this needs to be rounded to nearest whole number
			Bookmaker: outcome.BookmakerName,
			CashReturn: CashReturn }
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
		if bet.CashReturn <= cashInput {
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
