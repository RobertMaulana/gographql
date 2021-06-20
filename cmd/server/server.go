package main

import (
	"github.com/glyphack/graphlq-golang/config"
	"github.com/glyphack/graphlq-golang/graph"
	"github.com/glyphack/graphlq-golang/graph/generated"
	"github.com/glyphack/graphlq-golang/internal/auth"
	_ "github.com/glyphack/graphlq-golang/internal/auth"
	database "github.com/glyphack/graphlq-golang/internal/pkg/db/postgre"
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

	server := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	router := http.NewServeMux()
	router.Handle("/", playground.Handler("GraphQL playground", "/query"))
	router.Handle("/query", auth.Middleware(server))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}

/*
TODO:
1. logger query & mutation
2. rate limiter
3.
*/
