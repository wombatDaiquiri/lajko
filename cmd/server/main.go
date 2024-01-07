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
	// dsn := "host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai"
	// db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	// if err != nil {
	//	panic("failed to connect database")
	// }

	opts := []graphql.SchemaOpt{graphql.UseFieldResolvers()}
	schema := graphql.MustParseSchema(graphqlSchema, resolver.New(), opts...)

	r := chi.NewRouter()
	r.Use(
		middleware.Logger,
		cors.AllowAll().Handler,
	)

	r.Handle("/graphql", &relay.Handler{Schema: schema})
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("xd")
		w.Write([]byte("welcome"))
	})

	fmt.Println("listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
