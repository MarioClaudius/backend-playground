package main

import (
	// gql "backend/go/graphql"
	// gqlDB "backend/go/graphql/repository"

	"backend/go/rest"
	"backend/go/rest/handler"
	repo "backend/go/rest/repository/impl"
	uc "backend/go/rest/usecase/impl"
	"database/sql"
	"fmt"
	"log"
	"os"

	// "log"
	// "net/http"

	_ "github.com/go-sql-driver/mysql"
	// "github.com/graphql-go/graphql"
	// "github.com/graphql-go/handler"
	"github.com/joho/godotenv"
)

func main() {
	// get from env
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	// server := http.Server{
	// 	Addr:    "localhost:8080",
	// 	Handler: rest.GetHandler(),
	// }

	// err := server.ListenAndServe()
	// if err != nil {
	// 	log.Fatalln(err)
	// 	panic(err)
	// }
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", os.Getenv("MYSQL_USER"), os.Getenv("MYSQL_PASSWORD"), os.Getenv("MYSQL_HOST"), os.Getenv("MYSQL_PORT"), os.Getenv("MYSQL_DBNAME"))
	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	// Init Repo
	blogRepo := repo.NewBlogRepository(db)

	// Init Usecase
	blogUC := uc.NewBlogUC(blogRepo)

	// Init Handler
	h := handler.NewHandlers(blogUC)

	server := rest.NewAPIServer(":8081", h)
	server.Run()

	// log.Fatal(http.ListenAndServe(":8080", router))

	// ********************* GraphQL ***********************
	// gqlDB.InitDB()
	// gql.InitGQLFields()

	// schema, err := graphql.NewSchema(
	// 	graphql.SchemaConfig{
	// 		Query:    gql.QueryType(),
	// 		Mutation: gql.MutationType(),
	// 	},
	// )
	// if err != nil {
	// 	log.Fatalf("failed to create new schema, error: %v", err)
	// }

	// handler := handler.New(&handler.Config{
	// 	Schema: &schema,
	// 	Pretty: true,
	// })

	// http.Handle("/graphql", handler)
	// http.ListenAndServe(":8080", nil)
}
