package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/adiatma/moviedb-golang-graphql/src/handler"
	"github.com/friendsofgo/graphiql"
	"github.com/joho/godotenv"
)

func main() {
	graphiqlHandler, _ := graphiql.NewGraphiqlHandler("/graphql")

	http.HandleFunc("/graphql", handler.GraphqlHandler)
	http.Handle("/graphiql", graphiqlHandler)

	godotenvError := godotenv.Load()

	if godotenvError != nil {
		panic(godotenvError)
	}

	port := os.Getenv("PORT")
	host := os.Getenv("HOST")

	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", host, port), nil))
}
