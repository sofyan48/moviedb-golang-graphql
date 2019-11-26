package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/adiatma/moviedb-golang-graphql/src/config"
	"github.com/adiatma/moviedb-golang-graphql/src/handler"
	"github.com/friendsofgo/graphiql"
)

func main() {
	graphiqlHandler, _ := graphiql.NewGraphiqlHandler("/graphql")

	http.HandleFunc("/graphql", handler.GraphqlHandler)
	http.Handle("/graphiql", graphiqlHandler)

	// godotenvError := godotenv.Load()

	// if godotenvError != nil {
	// 	panic(godotenvError)
	// }

	enviroment := flag.String("e", "development", "")
	flag.Usage = func() {
		fmt.Println("Usage: server -e {mode}")
		os.Exit(1)
	}
	flag.Parse()
	config.Init(*enviroment)

	port := os.Getenv("PORT")
	host := os.Getenv("HOST")

	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", host, port), nil))
}
