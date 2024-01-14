package main

import (
	_ "embed"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"

	"github.com/wombatDaiquiri/lajko/resolver"
)

//go:embed schema.graphqls
var graphqlSchema string

func main() {
	opts := []graphql.SchemaOpt{graphql.UseFieldResolvers()}
	schema := graphql.MustParseSchema(graphqlSchema, resolver.New(), opts...)

	r := chi.NewRouter()
	r.Use(
		middleware.Logger,
		cors.AllowAll().Handler,
	)
	r.Handle("/graphql", &relay.Handler{Schema: schema})

	fmt.Println("listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
