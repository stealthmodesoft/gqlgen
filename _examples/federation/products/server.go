//go:generate go run ../../../testdata/gqlgen.go
package main

import (
	"log"
	"net/http"
	"os"

	"github.com/stealthmodesoft/gqlgen/_examples/federation/products/graph"
	"github.com/stealthmodesoft/gqlgen/_examples/federation/products/graph/generated"
	"github.com/stealthmodesoft/gqlgen/graphql/handler"
	"github.com/stealthmodesoft/gqlgen/graphql/handler/debug"
	"github.com/stealthmodesoft/gqlgen/graphql/playground"
)

const defaultPort = "4002"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))
	srv.Use(&debug.Tracer{})

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
