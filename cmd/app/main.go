package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/will3g/codefolio-backend/config"
	"github.com/will3g/codefolio-backend/handlers"
)

func main() {
	cfg := config.Load()

	db, err := config.ConnectDB(cfg)
	if err != nil {
		log.Fatalf("Error connecting to DB: %v", err)
	}

	_, article_index_err := config.CreateIndexes(db, "articles", "path")
	_, portfolio_index_err := config.CreateIndexes(db, "portfolio", "path")
	if portfolio_index_err != nil || article_index_err != nil {
		log.Fatalf("Error during creating index to path DB: %v", err)
	}

	router := mux.NewRouter()

	// router.HandleFunc("/portfolio/", handlers.GetPortfolio(db)).Methods(http.MethodGet)
	// router.HandleFunc("/portfolio/", handlers.CreatePortfolio(db)).Methods(http.MethodPost)
	// router.HandleFunc("/portfolio/{id}", handlers.UpdatePortfolio(db)).Methods(http.MethodPut)
	// router.HandleFunc("/portfolio/{id}", handlers.DeletePortfolio(db)).Methods(http.MethodDelete)

	router.HandleFunc("/articles/", handlers.GetArticles(db)).Methods(http.MethodGet)
	router.HandleFunc("/article/", handlers.GetArticle(db)).Methods(http.MethodGet)
	router.HandleFunc("/article/", handlers.UpdateArticle(db)).Methods(http.MethodPut)
	router.HandleFunc("/article/", handlers.CreateArticle(db)).Methods(http.MethodPost)
	router.HandleFunc("/article/", handlers.DeleteArticle(db)).Methods(http.MethodDelete)

	log.Printf("Server listening on port %s", "3030")
	if err := http.ListenAndServe(":"+"3030", router); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
