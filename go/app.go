package main

import (
	gql "backend/go/graphql"
	// "backend/go/rest"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
	"github.com/joho/godotenv"
)

func main() {
	// server := http.Server{
	// 	Addr:    "localhost:8080",
	// 	Handler: rest.GetHandler(),
	// }

	// err := server.ListenAndServe()
	// if err != nil {
	// 	log.Fatalln(err)
	// 	panic(err)
	// }

	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	gql.InitDB()
	blogType := gql.CreateBlogType()

	schema, err := graphql.NewSchema(
		graphql.SchemaConfig{
			Query: gql.QueryType(blogType),
		},
	)
	if err != nil {
		log.Fatalf("failed to create new schema, error: %v", err)
	}

	handler := handler.New(&handler.Config{
		Schema: &schema,
		Pretty: true,
	})

	http.Handle("/graphql", handler)
	http.ListenAndServe(":8080", nil)
}
