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

func MutationType() *graphql.Object {
	return graphql.NewObject(
		graphql.ObjectConfig{
			Name: "Mutation",
			Fields: graphql.Fields{
				"create_blog": CreateBlogField,
			},
		},
	)
}
