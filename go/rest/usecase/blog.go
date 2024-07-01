package usecase

import (
	"backend/go/rest/model/response"
	"context"
)

type BlogUsecase interface {
	GetBlogByID(ctx context.Context, id int64) (resp *response.BlogDetailResponse, statusCode int, err error)
}
