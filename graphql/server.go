// TODO(minkezhang): Delete package.
package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/minkezhang/truffle/api/graphql/model"
	"github.com/minkezhang/truffle/client/mal"
	"github.com/minkezhang/truffle/database"
	"github.com/minkezhang/truffle/graphql/resolver"

	graph "github.com/minkezhang/truffle/api/graphql"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &resolver.Resolver{
		DB: &resolver.DB{
			Entry: database.NewEntry(),
			APIData: map[model.APIType]*database.APIData{
				model.APITypeAPIMal: database.NewAPIData(
					mal.NewManga(
						mal.O{
							ClientID: mal.CLIENT_ID,
						},
					),
				),
			},
		},
	}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
