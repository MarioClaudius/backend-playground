package graphql

import (
	"backend/go/graphql/entity"
	"backend/go/graphql/handler"

	"github.com/graphql-go/graphql"
)

var GetBlogListField *graphql.Field
var GetBlogByIDField *graphql.Field
var CreateBlogField *graphql.Field

func InitGQLFields() {
	GetBlogListField = &graphql.Field{
		Type: graphql.NewList(entity.GetBlogType()), // it means the result will be in form of blog list
		Args: graphql.FieldConfigArgument{
			"limit": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
			"offset": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			return handler.GetBlogList(p)
		},
	}

	GetBlogByIDField = &graphql.Field{
		Type: entity.GetBlogType(),
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			return handler.GetBlogByID(p)
		},
	}

	CreateBlogField = &graphql.Field{
		Type:        entity.GetBlogType(),
		Description: "Create a new Blog",
		Args: graphql.FieldConfigArgument{
			"title": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"content": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			return handler.CreateBlog(p)
		},
	}
}
