package impl

import (
	"backend/go/rest/entity"
	"backend/go/rest/repository"
	"context"
	"database/sql"
	"errors"
)

type BlogRepository struct {
	db *sql.DB
}

func NewBlogRepository(db *sql.DB) repository.BlogRepository {
	return &BlogRepository{
		db: db,
	}
}

func (r *BlogRepository) GetBlogByID(ctx context.Context, id int64) (*entity.Blog, error) {
	var blog entity.Blog
	query := "SELECT id, title, content FROM blogs WHERE id = ?"
	err := r.db.QueryRowContext(ctx, query, id).Scan(&blog.ID, &blog.Title, &blog.Content)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("not found")
		}
		return nil, err
	}

	return &blog, nil
}
