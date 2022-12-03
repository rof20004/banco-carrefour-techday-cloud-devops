package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"github.com/google/uuid"
)

type CreateCommentRequest struct {
	Id      string `json:"id"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	Comment string `json:"comment"`
}

var (
	db  *sql.DB
	err error
)

func createDbConnection() (*sql.DB, error) {
	db, err = sql.Open("", "")
	if err != nil {
		log.Fatalln(err)
	}

	if err = db.Ping(); err != nil {
		return db, err
	}

	return db, err
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

	if err := saveComment(&payload); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	responseData, err := json.Marshal(&payload)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(responseData)
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "OPTIONS, POST, GET, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Content-Type, Accept, Authorization")
}

func closeRequestBody(r *http.Request) {
	_ = r.Body.Close()
}

func saveComment(comment *CreateCommentRequest) error {
	db, err := createDbConnection()
	if err != nil {
		return err
	}

	comment.Id = uuid.NewString()

	if _, err := db.Exec("INSERT INTO mensagens(id, name, email, comment) VALUES(?, ?, ?, ?)", comment.Id, comment.Name, comment.Email, comment.Comment); err != nil {
		return err
	}

	return nil
}
