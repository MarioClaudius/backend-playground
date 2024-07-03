package repository

import (
	"backend/go/rest/entity"
	"backend/go/rest/model/request"
	"context"
)

type BlogRepository interface {
	GetBlogByID(ctx context.Context, id int64) (*entity.Blog, error)
	GetBlogList(ctx context.Context, filter request.BlogListFilter) ([]entity.Blog, error)
	InsertBlog(ctx context.Context, data entity.Blog) (int64, error)
}
