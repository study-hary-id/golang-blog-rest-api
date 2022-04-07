package main

import (
	"log"
	"net/http"
)

type articleHandler struct {}

func (h *articleHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {}

func main () {
	const PORT = ":8000"
	mux := http.NewServeMux()
	mux.Handle("/articles", &articleHandler{})
	log.Printf("Server listening at http://localhost%v\n", PORT)
	log.Fatal(http.ListenAndServe(PORT, mux))
}
