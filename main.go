package main

import (
	dg "desafio-graph/graph"
	"desafio-graph/src/libs"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/handler"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	http.Handle("/", handler.Playground("GraphQL playground", "/query"))

	resolvers := dg.NewResolver(libs.CrawlerImpl{})
	executableShema := dg.NewExecutableSchema(dg.Config{Resolvers: resolvers})
	http.Handle("/query", handler.GraphQL(executableShema))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
