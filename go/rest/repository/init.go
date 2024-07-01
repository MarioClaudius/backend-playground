package repository

import (
	"backend/go/rest/entity"
	"context"
)

type BlogRepository interface {
	GetBlogByID(ctx context.Context, id int64) (*entity.Blog, error)
}
