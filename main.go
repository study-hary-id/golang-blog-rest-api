package main

import (
	"encoding/json"
	"log"
	"net/http"
)

// articleHandler is a Handler that handles articles operation.
type articleHandler struct{}

// List responses the list of all available articles.
func (h *articleHandler) List(w http.ResponseWriter) {
	jsonBytes, err := json.Marshal(articles)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(jsonBytes)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// ServeHTTP is a default handler to given endpoint pattern from ServeMux.
func (h *articleHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	switch {
	case r.Method == http.MethodGet && listArticleRe.MatchString(r.URL.Path):
		h.List(w)
		return
	}
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
