package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/friendsofgo/graphiql"
	"github.com/graphql-go/graphql"
)

func main() {
	fields := graphql.Fields{
		"hello": &graphql.Field{
			Type: graphql.String,
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				return "world", nil
			},
		},
	}

	rootQuery := graphql.ObjectConfig{Name: "RootQuery", Fields: fields}

	schemaConfig := graphql.SchemaConfig{Query: graphql.NewObject(rootQuery)}

	schema, _ := graphql.NewSchema(schemaConfig)

	http.HandleFunc("/graphql", func(w http.ResponseWriter, r *http.Request) {
		body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))

		if err != nil {
			log.Fatalln("Error", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		var apolloQuery map[string]interface{}

		json.Unmarshal(body, &apolloQuery)

		query := apolloQuery["query"]

		params := graphql.Params{
			Schema:        schema,
			RequestString: query.(string),
		}

		result := graphql.Do(params)

		if len(result.Errors) > 0 {
			fmt.Printf("Error %+v", result.Errors)
		}

		json.NewEncoder(w).Encode(result)
	})

	graphiqlHandler, err := graphiql.NewGraphiqlHandler("/graphql")

	if err != nil {
		panic(err)
	}

	http.Handle("/graphiql", graphiqlHandler)

	http.ListenAndServe(":8080", nil)
}
