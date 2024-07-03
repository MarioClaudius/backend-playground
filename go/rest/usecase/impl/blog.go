package impl

import (
	"backend/go/rest/entity"
	"backend/go/rest/model/request"
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

func (b *blogUsecase) GetBlogList(ctx context.Context, filter request.BlogListFilter) (resp []*response.BlogDetailResponse, statusCode int, err error) {
	blogDBList, err := b.blogRepo.GetBlogList(ctx, filter)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	for _, blogDB := range blogDBList {
		resp = append(resp, &response.BlogDetailResponse{
			ID:      blogDB.ID,
			Title:   blogDB.Title,
			Content: blogDB.Content,
		})
	}

	return resp, http.StatusOK, nil
}

func (b *blogUsecase) CreateBlog(ctx context.Context, req request.CreateBlogRequest) (statusCode int, err error) {
	data := entity.Blog{
		Title:   req.Title,
		Content: req.Content,
	}

	_, err = b.blogRepo.InsertBlog(ctx, data)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, nil
}
