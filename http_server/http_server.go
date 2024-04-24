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
	w.Header().Set("Content-Type", "application/json")
	
	err := wrangling.FetchAndStoreData()
	if err != nil {
		json.NewEncoder(w).Encode(Response{ Err: err.Error() })
	}

	json.NewEncoder(w).Encode(Response{ Message: "Data has been refreshed"})
}

func StartHttpServer() {
	fmt.Println("Starting http server")
	http.HandleFunc("/", renderHtml)
	http.HandleFunc("/refresh-bets", refreshBetData)
	http.ListenAndServe(":8080", nil)
}
