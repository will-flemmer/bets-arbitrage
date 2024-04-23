package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Sport struct {
	Key          string `json:"key"`
	Group        string `json:"group"`
	Title        string `json:"title"`
	Description  string `json:"description"`
	Active       bool   `json:"active"`
	HasOutrights bool   `json:"has_outrights"`
}

func GetJson[T any](url string, target *[]T) error {
	r, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer r.Body.Close()

	used := r.Header.Get("x-requests-used")
	remaining := r.Header.Get("x-requests-remaining")
	fmt.Println("Used", used, "requests")
	fmt.Println(remaining, "reamaining requests")

	body, err := ioutil.ReadAll(r.Body)
	return json.Unmarshal(body, target)
}

func CheckRemainingRequests(apiToken string) {
	endpoint := fmt.Sprintf("https://api.the-odds-api.com/v4/sports/?apiKey=%s", apiToken)
	r, err := http.Get(endpoint)
	if err != nil {
		panic(err)
	}
	used := r.Header.Get("x-requests-used")
	remaining := r.Header.Get("x-requests-remaining")
	fmt.Println("Used", used, "requests")
	fmt.Println(remaining, "reamaining requests")
}

