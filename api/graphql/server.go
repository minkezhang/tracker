package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/minkezhang/truffle/api/graphql/resolver"

	graph "github.com/minkezhang/truffle/api/graphql/generated"
)

const defaultPort = "8080"

// Example:
//
//	mutation {
//	  entry(input: {
//	    corpus: CORPUS_ANIME,
//	    queued: false,
//	    titles: [{
//	        language: "en",
//	        title: "Neon Genesis Evangelion"
//	    }],
//	    providers: [
//	      PROVIDER_NETFLIX,
//	    ],
//	    tags: [
//	      "mechs", "psychological",
//	    ]
//	    aux: {
//	      studios: ["Gainax", "Tatsunoko"]
//	    }
//	    sources: [
//	      {
//	        api: API_MAL,
//	        id: "anime/30"
//	      }
//	    ]
//	  }) {
//	    id
//	    corpus
//	    queued
//	    metadata {
//	      truffle {
//	        cached
//	        api
//	        id
//	        titles {
//	          language
//	          title
//	        }
//	        aux {
//	          ... on AuxAnime {
//	            studios
//	          }
//	        },
//	        providers,
//	        tags,
//	      }
//	      sources {
//	      	id
//	        cached
//	      }
//	    }
//	  }
//	}
func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &resolver.Resolver{}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
