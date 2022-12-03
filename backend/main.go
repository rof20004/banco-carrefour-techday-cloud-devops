package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type CreateCommentRequest struct {
	Name    string `json:"name"`
	Email   string `json:"email"`
	Comment string `json:"comment"`
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/comments", createComment)

	log.Println("Server started at port 8080")

	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Println(err)
	}
}

func createComment(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodOptions {
		return
	}

	enableCors(&w)

	defer closeRequestBody(r)

	var payload CreateCommentRequest

	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "OPTIONS, POST, GET, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Content-Type, Accept, Authorization")
}

func closeRequestBody(r *http.Request) {
	_ = r.Body.Close()
}
