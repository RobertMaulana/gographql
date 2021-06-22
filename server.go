package main

import (
	"github.com/RobertMaulana/graphql-go/config"
	"github.com/RobertMaulana/graphql-go/graph"
	"github.com/RobertMaulana/graphql-go/graph/generated"
	"github.com/RobertMaulana/graphql-go/internal/auth"
	database "github.com/RobertMaulana/graphql-go/internal/pkg/db/postgre"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

func init() {
	_ = config.LoadENV()
	database.InitDB()
	database.Migrate()
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	server := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{
		Resolvers: &graph.Resolver{},
	}))

	router := http.NewServeMux()
	router.Handle("/", playground.Handler("GraphQL playground", "/query"))
	router.Handle("/query", auth.Middleware(server))

	infoLog.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	errorLog.Fatal(http.ListenAndServe(":"+port, router))
}