package handler

import (
	"backend/go/graphql/entity"
	blogRepo "backend/go/graphql/repository"
	"errors"

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

func CreateBlog(p graphql.ResolveParams) (interface{}, error) {
	title, titleValid := p.Args["title"].(string)
	content, contentValid := p.Args["content"].(string)
	if !titleValid || !contentValid {
		return nil, errors.New("invalid argument")
	}

	model := entity.Blog{
		Title:   title,
		Content: content,
	}

	return blogRepo.InsertBlog(model)
}
