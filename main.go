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

	executableShema := dg.NewExecutableSchema(dg.NewResolver(libs.CrawlerImpl{}, libs.MonetaryImpl{}))
	http.Handle("/query", handler.GraphQL(executableShema))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
