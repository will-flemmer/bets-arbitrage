package httpServer

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"scraping/wrangling"
)


type FrontendData struct {
	Message string
}

func renderHtml(w http.ResponseWriter, _ *http.Request) {
	data := FrontendData{Message: "Lets find some bets!"}
	template, err := template.ParseFiles("http_server/main.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = template.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}


type Response struct {
	Err string
	Message string
}

func refreshBetData(w http.ResponseWriter, req *http.Request) {
	err := wrangling.FetchAndStoreData()
	
	w.Header().Set("Content-Type", "application/json")

	if err != nil {
		println(err.Error())
		payload, err := json.Marshal(Response{ Err: err.Error() }) 
		if err != nil {
			panic(err)
		}
		json.NewEncoder(w).Encode(payload)
	}

	payload, err := json.Marshal(Response{ Message: "Data has been refreshed"})
	if err != nil {
		panic(err)
	}

	json.NewEncoder(w).Encode(payload)
}

func StartHttpServer() {
	fmt.Println("Starting http server")
	http.HandleFunc("/", renderHtml)
	http.HandleFunc("/refresh-bets", refreshBetData)
	http.ListenAndServe(":8080", nil)
}
