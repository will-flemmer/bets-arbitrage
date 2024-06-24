package utils

import (
	"errors"
	"log"
	"os"

	"github.com/spf13/viper"
	"gopkg.in/yaml.v2"
)

type HasYaml interface {
	FileName() string
}

type Env struct {
	API_TOKEN string `yaml:"API_TOKEN"`
}

func (e *Env) FileName() string {
	return "env.yaml"
}

func LoadEnv() Env {
	env := &Env{}
	readYaml(env)
	return *env
}

type Sports struct {
	Soccer []string `yaml:"soccer"`
}

func (s *Sports) FileName() string {
	return "sports.yaml"
}

func LoadSports() Sports {
	s := &Sports{}
	readYaml(s)
	return *s
}

type DatabaseEnvConfig struct {
	Database string `yaml:"database"`
  Host string `yaml:"host"`
  Username string `yaml:"username"`
	Password string `yaml:"password"`
}

type DatabaseConfig struct {
	Development DatabaseEnvConfig `yaml:"development"`
	Test DatabaseEnvConfig `yaml:"test"`
}

func (dbConfig *DatabaseConfig) FileName() string {
	return "config/database.yaml"
}

func LoadDataBaseConfig() (DatabaseEnvConfig, error) {
	databaseConfig := &DatabaseConfig{}
	readYaml(databaseConfig)
	env := viper.Get("ENV")
	if env == "development" {
		return databaseConfig.Development, nil
	}
	if env == "test" {
		return databaseConfig.Test, nil
	}
	return databaseConfig.Development, errors.New("invalid database config")
}

func readYaml[T HasYaml](outputStruct T) {
	file := outputStruct.FileName()
	yamlFile, err := os.ReadFile(file)
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, outputStruct)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}
}