package entity

import (
	"github.com/graphql-go/graphql"
)

var blogType *graphql.Object

type Blog struct {
	ID      int64  `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

// need to be a singleton so the GQL field doesnt identify that there're more than 1 type with the same name
func GetBlogType() *graphql.Object {
	if blogType == nil {
		blogType = graphql.NewObject(
			graphql.ObjectConfig{
				Name: "Blog",
				Fields: graphql.Fields{
					"id": &graphql.Field{
						Type: graphql.Int,
					},
					"title": &graphql.Field{ // key in this schema must matching with with the field name in struct but no uppercase
						Type: graphql.String,
					},
					"content": &graphql.Field{
						Type: graphql.String,
					},
				},
			},
		)
	}
	return blogType
}
