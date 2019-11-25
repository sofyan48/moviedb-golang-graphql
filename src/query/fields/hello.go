package fields

import (
	gql "github.com/graphql-go/graphql"
)

// Hello fields
var Hello = gql.Fields{
	"hello": &gql.Field{
		Type: gql.String,
		Resolve: func(params gql.ResolveParams) (interface{}, error) {
			return "world", nil
		},
	},
}
