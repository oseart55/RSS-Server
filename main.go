package main

import (
	"database/sql"
	"log"
	"main/internal/database"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"

	_ "github.com/lib/pq"
)

type apiConfig struct {
	DB *database.Queries
}

func main() {
	godotenv.Load()
	portString := os.Getenv("PORT")
	DB_URL := os.Getenv("DB_URL")

	if portString == "" {
		log.Fatal("Could not load PORT from ENV")
	}
	if DB_URL == "" {
		log.Fatal("Could not load PORT from ENV")
	}

	conn, err := sql.Open("postgres", DB_URL)
	if err != nil {
		log.Fatal("Error connecting to Database")
	}
	apiCfg := apiConfig{
		DB: database.New(conn),
	}

	router := chi.NewRouter()
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"Get", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	v1Router := chi.NewRouter()

	v1Router.Get("/health", handleReady)
	v1Router.Get("/error", handleError)

	v1Router.Post("/users", apiCfg.handleCreateUser)
	v1Router.Get("/users", apiCfg.middlewareAuth(apiCfg.handleGetUser))

	v1Router.Post("/feeds", apiCfg.middlewareAuth(apiCfg.handleCreateFeed))
	v1Router.Get("/feeds", apiCfg.middlewareAuth(apiCfg.handleGetFeeds))

	router.Mount("/v1", v1Router)
	router.Handle("/", http.FileServer(http.Dir("./static")))
	srv := &http.Server{
		Handler: router,
		Addr:    ":" + portString,
	}

	log.Printf("Server Running on Port: %s", portString)
	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
