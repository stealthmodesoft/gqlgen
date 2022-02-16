package main

import (
	"log"
	"net/http"

	"github.com/stealthmodesoft/gqlgen/_examples/scalars"
	"github.com/stealthmodesoft/gqlgen/graphql/handler"
	"github.com/stealthmodesoft/gqlgen/graphql/playground"
)

func main() {
	http.Handle("/", playground.Handler("Starwars", "/query"))
	http.Handle("/query", handler.NewDefaultServer(scalars.NewExecutableSchema(scalars.Config{Resolvers: &scalars.Resolver{}})))

	log.Fatal(http.ListenAndServe(":8084", nil))
}
