package handler

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/adiatma/moviedb-golang-graphql/src/schema"
	gql "github.com/graphql-go/graphql"
)

// QueryInterface map
var QueryInterface map[string]interface{}

// GraphqlHandler methods
func GraphqlHandler(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	json.Unmarshal(body, &QueryInterface)
	schema, _ := schema.Schema()
	query := QueryInterface["query"]
	result := gql.Do(gql.Params{
		Schema:        schema,
		RequestString: query.(string),
	})
	json.NewEncoder(w).Encode(result)
}
