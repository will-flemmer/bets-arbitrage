package wrangling

import (
	"fmt"
	"log"
	"scraping/utils"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Outcome struct {
	ID       uint `json:"-"`
	MarketID uint
	Market   Market
	Name     string  `json:"name"`
	Price    float32 `json:"price"`
}

type Market struct {
	ID          uint `json:"-"`
	BookmakerID uint
	Bookmaker   Bookmaker
	Key         string    `json:"key"`
	LastUpdate  string    `json:"last_update"`
	Outcomes    []Outcome `json:"outcomes"`
}

type Bookmaker struct {
	ID         uint `json:"-"`
	FixtureID  uint
	Fixture    Fixture
	Key        string   `json:"key"`
	LastUpdate string   `json:"last_update"`
	Markets    []Market `json:"markets" gorm:"foreignKey:BookmakerID"`
	Title      string   `json:"title"`
}

type Fixture struct {
	ID         uint        `json:"-"`
	CreatedAt  time.Time   `json:"-"`
	AwayTeam   string      `json:"away_team" gorm:"uniqueIndex"`
	Bookmakers []Bookmaker `json:"bookmakers" gorm:"foreignKey:FixtureID"`
}

type BestOdds struct {
	TeamName      string  `json:"teamName"`
	Odds          float32 `json:"odds"`
	BookmakerName string  `json:"bookmakerName"`
}

const (
	DB_NAME      = "dev.db"
	TEST_DB_NAME = "test.db"
)

var GlobalDB *gorm.DB

func LoadDb() *gorm.DB {
	dbConfig, err := utils.LoadDataBaseConfig()
	if err != nil {
		log.Panic(err)
	}

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=9920 sslmode=disable TimeZone=Asia/Shanghai",
		dbConfig.Host, dbConfig.Username, dbConfig.Password, dbConfig.Database,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Panic("Could not open db connection")
	}

	db.AutoMigrate(&Outcome{}, &Fixture{}, &Bookmaker{}, &Market{})
	GlobalDB = db

	return db
}
