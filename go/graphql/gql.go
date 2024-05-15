package graphql

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/graphql-go/graphql"
)

type Blog struct {
	ID      int64  `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

var db *sql.DB

func InitDB() {
	var err error
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", os.Getenv("MYSQL_USER"), os.Getenv("MYSQL_PASSWORD"), os.Getenv("MYSQL_HOST"), os.Getenv("MYSQL_PORT"), os.Getenv("MYSQL_DBNAME"))
	db, err = sql.Open("mysql", dataSourceName)
	if err != nil {
		log.Fatal(err)
	}
}

func CreateBlogType() *graphql.Object {
	return graphql.NewObject(
		graphql.ObjectConfig{
			Name: "Blog",
			Fields: graphql.Fields{
				"id": &graphql.Field{
					Type: graphql.Int,
				},
				"title": &graphql.Field{
					Type: graphql.String,
				},
				"content": &graphql.Field{
					Type: graphql.String,
				},
			},
		},
	)
}

func QueryType(blogType *graphql.Object) *graphql.Object {
	return graphql.NewObject(
		graphql.ObjectConfig{
			Name: "Query",
			Fields: graphql.Fields{
				"blogs": &graphql.Field{
					Type: graphql.NewList(blogType),
					Args: graphql.FieldConfigArgument{
						"limit": &graphql.ArgumentConfig{
							Type: graphql.Int,
						},
						"offset": &graphql.ArgumentConfig{
							Type: graphql.Int,
						},
					},
					Resolve: func(p graphql.ResolveParams) (interface{}, error) {
						limit, _ := p.Args["limit"].(int)
						if limit <= 0 || limit > 20 {
							limit = 10
						}
						offset, _ := p.Args["offset"].(int)
						if offset < 0 {
							offset = 0
						}
						return getBlogs(limit, offset)
					},
				},
			},
		},
	)
}

func getBlogs(limit, offset int) ([]Blog, error) {
	var blogs []Blog
	query := fmt.Sprintf("SELECT id, title, content FROM blogs limit %d offset %d", limit, offset)
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var b Blog
		if err := rows.Scan(&b.ID, &b.Title, &b.Content); err != nil {
			return nil, err
		}
		blogs = append(blogs, b)
	}
	return blogs, nil
}

// func main() {
// 	initDB()

// 	blogType := createBlogType()

// 	schema, err := graphql.NewSchema(
// 		graphql.SchemaConfig{
// 			Query: queryType(blogType),
// 		},
// 	)
// 	if err != nil {
// 		log.Fatalf("failed to create new schema, error: %v", err)
// 	}

// 	handler := handler.New(&handler.Config{
// 		Schema: &schema,
// 		Pretty: true,
// 	})

// 	http.Handle("/graphql", handler)
// 	http.ListenAndServe(":8080", nil)
// }
