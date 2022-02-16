package main

import (
	"log"
	"net/http"

	todo "github.com/stealthmodesoft/gqlgen/_examples/config"
	"github.com/stealthmodesoft/gqlgen/graphql/handler"
	"github.com/stealthmodesoft/gqlgen/graphql/playground"
)

func main() {
	http.Handle("/", playground.Handler("Todo", "/query"))
	http.Handle("/query", handler.NewDefaultServer(
		todo.NewExecutableSchema(todo.New()),
	))
	log.Fatal(http.ListenAndServe(":8081", nil))
}
