package utils

import (
	"log"
	"os"

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