package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi"
	"github.com/joho/godotenv"
	"github.com/pranotobudi/go-graphql-hackernews/config"
	"github.com/pranotobudi/go-graphql-hackernews/database"
	"github.com/pranotobudi/go-graphql-hackernews/graph"
	"github.com/pranotobudi/go-graphql-hackernews/graph/generated"
	"github.com/pranotobudi/go-graphql-hackernews/internal/auth"
)

// reference: https://www.howtographql.com/graphql-go/0-introduction/

const defaultPort = "8080"

func main() {
	if os.Getenv("APP_ENV") != "production" {
		// executed in development only,
		//for production set those on production environment settings

		// load local env variables to os
		err := godotenv.Load(".env")
		if err != nil {
			log.Println("failed to load .env file")
		}
	}

	db := database.InitDB()
	db.MigrateDB("./database/init.sql")

	router := chi.NewRouter()
	router.Use(auth.Middleware())

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	router.Handle("/", playground.Handler("GraphQL playground", "/query"))
	router.Handle("/query", srv)

	config := config.AppConfig()
	log.Printf("connect to http://localhost:%s/ for GraphQL playground", config.Port)
	log.Fatal(http.ListenAndServe(":"+config.Port, router))
}
