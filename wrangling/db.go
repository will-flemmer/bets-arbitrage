package wrangling

import (
	"log"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Outcome struct {
	ID uint `json:"-"`
	MarketID uint
	Market Market
	Name  string  `json:"name"`
	Price float32 `json:"price"`
}

type Market struct {
	ID uint `json:"-"`
	BookmakerID uint
	Bookmaker Bookmaker
	Key        string `json:"key"`
	LastUpdate string `json:"last_update"`
	Outcomes   []Outcome `json:"outcomes"`
}

type Bookmaker struct {
	ID uint `json:"-"`
	FixtureID uint
	Fixture Fixture
	Key        string `json:"key"`
	LastUpdate string `json:"last_update"`
	Markets    []Market `json:"markets" gorm:"foreignKey:BookmakerID"`
	Title string `json:"title"`
}

type Fixture struct {
	ID uint `json:"-"`
	CreatedAt time.Time `json:"-"`
	AwayTeam string `json:"away_team" gorm:"uniqueIndex"`
	Bookmakers []Bookmaker `json:"bookmakers" gorm:"foreignKey:FixtureID"`
}

type BestOdds struct {
	TeamName      string `json:"teamName"`
	Odds          float32 `json:"odds"`
	BookmakerName string `json:"bookmakerName"`
}

const (
	DB_NAME = "dev.db"
)


func LoadDb() *gorm.DB {
	db, err := gorm.Open(sqlite.Open(DB_NAME), &gorm.Config{})
	if err != nil {
		log.Panic("Could not open db connection")
	}

	db.AutoMigrate(&Outcome{}, &Fixture{}, &Bookmaker{}, &Market{})

	return db
}
