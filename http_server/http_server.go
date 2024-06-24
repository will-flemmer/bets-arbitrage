package httpServer

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"scraping/bets"
	"scraping/graphqlApi"
	"scraping/wrangling"

	"github.com/graphql-go/handler"
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

type AnalysisResponse struct {
	ProfitableBets []bets.ProfitableBet `json:"profitableBets"`
	// ProfitableBets [][]bets.Bet `json:"profitableBets"`
}

func refreshBetData(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	
	err := wrangling.FetchAndStoreData()
	if err != nil {
		json.NewEncoder(w).Encode(Response{ Err: err.Error() })
		return
	}

	json.NewEncoder(w).Encode(Response{ Message: "Data has been refreshed"})
}

func runAnalysis(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	
	profitableBets := bets.FindBets()
	json.NewEncoder(w).Encode(AnalysisResponse{ ProfitableBets: profitableBets })
}

func generateGraphqlShema(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	
	introspection := graphqlApi.GenerateSchema()
	json.NewEncoder(w).Encode(introspection)
}

func StartHttpServer() {
	fmt.Println("Starting http server")

	mux := http.NewServeMux()
	mux.HandleFunc("/", renderHtml)
	mux.HandleFunc("/refresh-bets", refreshBetData)
	mux.HandleFunc("/run-analysis", runAnalysis)
	
	mux.HandleFunc("/generate-schema", generateGraphqlShema)
	schema := graphqlApi.CreateSchema()
	graphqlHandle := handler.New(&handler.Config{
		Schema: &schema,
		Pretty: true,
		GraphiQL: true,
	})
	mux.Handle("/graphql", graphqlHandle)

	// must enable when using with react
	newMux := UseCorsMiddleWare(mux)
	http.ListenAndServe(":8080", newMux)
}
