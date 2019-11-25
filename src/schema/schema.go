package schema

import (
	"github.com/adiatma/moviedb-golang-graphql/src/query"
	gql "github.com/graphql-go/graphql"
)

// Schema config
func Schema() (gql.Schema, error) {
	schema, error := gql.NewSchema(gql.SchemaConfig{
		Query: gql.NewObject(query.Query),
	})

	return schema, error
}
