package handler

import (
	blogRepo "backend/go/graphql/repository"

	"github.com/graphql-go/graphql"
)

func GetBlogList(p graphql.ResolveParams) (interface{}, error) {
	limit, _ := p.Args["limit"].(int)
	if limit <= 0 || limit > 20 {
		limit = 10
	}
	offset, _ := p.Args["offset"].(int)
	if offset < 0 {
		offset = 0
	}
	return blogRepo.GetBlogs(limit, offset)
}

func GetBlogByID(p graphql.ResolveParams) (interface{}, error) {
	blogID, _ := p.Args["id"].(int)
	return blogRepo.GetBlogByID(int64(blogID))
}
