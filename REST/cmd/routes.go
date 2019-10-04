package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"
	ra "github.com/shanehowearth/cloudblog/readarticle/integration/grpc/client/v1"
)

var bc = ra.BlogClient{}

func blogRoutes() *chi.Mux {
	router := chi.NewRouter()
	router.Get("/Articles", GetArticles)
	return router
}

// GetArticles -
func GetArticles(w http.ResponseWriter, r *http.Request) {
	al, err := bc.GetArticles()
	if err != nil {
		// log the error
		log.Printf("An error occurred with GetArticles, Error: %v", err)
		respondWithError(w, http.StatusInternalServerError, "An unexpected error has occurred, the issue has been reported to our engineers and will be looked into")
	}
	respondWithJSON(w, http.StatusOK, al)

}
