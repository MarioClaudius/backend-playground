package handler

import "backend/go/rest/usecase"

type Handlers struct {
	blogUC usecase.BlogUsecase
}

func NewHandlers(blogUC usecase.BlogUsecase) *Handlers {
	return &Handlers{
		blogUC: blogUC,
	}
}
