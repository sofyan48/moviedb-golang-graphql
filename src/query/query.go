package query

import (
	"github.com/adiatma/moviedb-golang-graphql/src/query/fields"
	gql "github.com/graphql-go/graphql"
)

// Query object
var Query = gql.ObjectConfig{
	Name:   "Query",
	Fields: fields.Hello,
}
