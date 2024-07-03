package impl

import (
	"backend/go/rest/entity"
	"backend/go/rest/model/request"
	"backend/go/rest/repository"
	"context"
	"database/sql"
	"errors"
	"fmt"
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

func (r *BlogRepository) GetBlogList(ctx context.Context, filter request.BlogListFilter) ([]entity.Blog, error) {
	query := "SELECT id, title, content FROM blogs "
	args := []interface{}{}

	if filter.Title != "" {
		query = fmt.Sprintf("%s %s", query, "WHERE title LIKE ?")
		args = append(args, "%"+filter.Title+"%")
	}

	if filter.Limit > 0 {
		query = fmt.Sprintf("%s %s", query, "LIMIT ?")
		args = append(args, filter.Limit)
	}

	if filter.Offset > 0 {
		query = fmt.Sprintf("%s %s", query, "OFFSET ?")
		args = append(args, filter.Offset)
	}

	rows, err := r.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var blogs []entity.Blog
	for rows.Next() {
		var blog entity.Blog
		if err := rows.Scan(&blog.ID, &blog.Title, &blog.Content); err != nil {
			return nil, err
		}
		blogs = append(blogs, blog)
	}

	return blogs, nil
}

func (r *BlogRepository) InsertBlog(ctx context.Context, data entity.Blog) (int64, error) {
	q := "INSERT INTO blogs (title, content) VALUES (?,?)"
	res, err := r.db.ExecContext(ctx, q, data.Title, data.Content)
	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, nil
}
