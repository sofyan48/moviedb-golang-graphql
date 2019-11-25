package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/friendsofgo/graphiql"
	"github.com/graphql-go/graphql"
	"github.com/joho/godotenv"
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
		body, bodyError := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))

		if bodyError != nil {
			log.Fatalln("Error", bodyError)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		var storageQuery map[string]interface{}

		json.Unmarshal(body, &storageQuery)

		query := storageQuery["query"]

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

	graphiqlHandler, graphiqlError := graphiql.NewGraphiqlHandler("/graphql")

	if graphiqlError != nil {
		panic(graphiqlError)
	}

	http.Handle("/graphiql", graphiqlHandler)

	godotenvError := godotenv.Load()

	if godotenvError != nil {
		panic(godotenvError)
	}

	port := os.Getenv("PORT")
	host := os.Getenv("HOST")

	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", host, port), nil))
}
