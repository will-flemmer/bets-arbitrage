package graphqlApi

import (
	"net/http"
	"net/http/httptest"
	"scraping/wrangling"
	"testing"

	"github.com/davecgh/go-spew/spew"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
)

func createGraphqlServer() (*httptest.ResponseRecorder, *http.ServeMux) {
	schema := CreateSchema()
	graphqlHandler := handler.New(&handler.Config{
		Schema: &schema,
		Pretty: false,
		GraphiQL: false,
	})

	mux := http.NewServeMux()
	mux.Handle("/graphql", graphqlHandler)
	rr := httptest.NewRecorder()
	return rr, mux
}

type GraphqlRequest struct {
	Query string `json:"string"`
	Variables map[string]interface{} `json:"variables,omitempty"`
}

func sendQuery(input GraphqlRequest) *graphql.Result {
	params := graphql.Params{
		RequestString: input.Query,
		Schema: CreateSchema(),
	}
	return graphql.Do(params)
}

func TestReturnsFixtures(t *testing.T) {
	db := wrangling.LoadDb()
	query := `{
		allfixtures(limit: 4) {
			awayTeam
			bookmakers {
				markets {
					key
				}
			}
		}
	}
	`

	var fixtures []wrangling.Fixture
	db.Find(&fixtures)
	spew.Dump(fixtures)

	variables := map[string]interface{}{
		"limit": 2,
	}
	result := sendQuery(GraphqlRequest{
		Query: query,
		Variables: variables,
	})

	if result.HasErrors() {
		t.Errorf("response has errors: %v", result.Errors)
	}

	// Check the response body
	expected := "Hello, World!"
	if "meow" != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", result.Data, expected)
	}
}
