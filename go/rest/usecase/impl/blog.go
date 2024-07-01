package impl

import (
	"backend/go/rest/model/response"
	"backend/go/rest/repository"
	"backend/go/rest/usecase"
	"context"
	"net/http"
)

type blogUsecase struct {
	blogRepo repository.BlogRepository
}

func NewBlogUC(repo repository.BlogRepository) usecase.BlogUsecase {
	return &blogUsecase{
		blogRepo: repo,
	}
}

func (b *blogUsecase) GetBlogByID(ctx context.Context, id int64) (resp *response.BlogDetailResponse, statusCode int, err error) {
	blogDB, err := b.blogRepo.GetBlogByID(ctx, id)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	resp = &response.BlogDetailResponse{
		ID:      blogDB.ID,
		Title:   blogDB.Title,
		Content: blogDB.Content,
	}

	return resp, http.StatusOK, nil
}
