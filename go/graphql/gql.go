package graphql

import (
	"github.com/graphql-go/graphql"
)

func QueryType() *graphql.Object {
	return graphql.NewObject(
		graphql.ObjectConfig{
			Name: "Query",
			Fields: graphql.Fields{
				"blogs": GetBlogListField,
				"blog":  GetBlogByIDField,
			},
		},
	)
}
