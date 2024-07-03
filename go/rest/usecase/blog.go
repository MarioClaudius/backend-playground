package usecase

import (
	"backend/go/rest/model/request"
	"backend/go/rest/model/response"
	"context"
)

type BlogUsecase interface {
	GetBlogByID(ctx context.Context, id int64) (resp *response.BlogDetailResponse, statusCode int, err error)
	GetBlogList(ctx context.Context, filter request.BlogListFilter) (resp []*response.BlogDetailResponse, statusCode int, err error)
	CreateBlog(ctx context.Context, req request.CreateBlogRequest) (statusCode int, err error)
}
