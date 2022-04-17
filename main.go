package main

import (
	"log"
	"net/http"
)

// articleHandler is a Handler that handles articles operation.
type articleHandler struct{}

// ServeHTTP is a default listener for each custom multiplexer,
// it handles every subtrees of given endpoint pattern.
func (h *articleHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	switch r.Method {
	case http.MethodGet:
		h.List(w, r)
		return
	}
}

// List gives all article resources.
func (h *articleHandler) List(w http.ResponseWriter, r *http.Request) {
	articles, err := getArticles()
	if err != nil {
		internalServerError(w, err)
	}
	responseJSON(w, r, articles, http.StatusOK)
}

func main() {
	// configurations
	const port = ":8000"
	mux := http.NewServeMux()

	// custom handlers
	mux.Handle("/articles/", &articleHandler{})

	log.Printf("Server listening at http://localhost%v\n", port)
	log.Fatal(http.ListenAndServe(port, mux))
}
