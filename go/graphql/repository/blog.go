package repository

import (
	"backend/go/graphql/entity"
	"database/sql"
	"fmt"
	"os"
)

var db *sql.DB

func InitDB() error {
	var err error
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", os.Getenv("MYSQL_USER"), os.Getenv("MYSQL_PASSWORD"), os.Getenv("MYSQL_HOST"), os.Getenv("MYSQL_PORT"), os.Getenv("MYSQL_DBNAME"))
	db, err = sql.Open("mysql", dataSourceName)
	return err
}

func GetBlogs(limit, offset int) ([]entity.Blog, error) {
	var blogs []entity.Blog
	query := fmt.Sprintf("SELECT id, title, content FROM blogs limit %d offset %d", limit, offset)
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var b entity.Blog
		if err := rows.Scan(&b.ID, &b.Title, &b.Content); err != nil {
			return nil, err
		}
		blogs = append(blogs, b)
	}
	return blogs, nil
}

func GetBlogByID(blogID int64) (*entity.Blog, error) {
	var blog entity.Blog
	query := fmt.Sprintf("SELECT id, title, content FROM blogs WHERE id = %d", blogID)
	err := db.QueryRow(query).Scan(&blog.ID, &blog.Title, &blog.Content)
	if err != nil {
		return nil, err
	}
	return &blog, nil
}
