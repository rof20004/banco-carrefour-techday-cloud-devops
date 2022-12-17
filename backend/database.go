package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var (
	dbConn   *sql.DB
	dbDriver = "mysql"
	dbHost   = os.Getenv("DB_HOST")
	dbPort   = os.Getenv("DB_PORT")
	dbUser   = os.Getenv("DB_USER")
	dbPass   = os.Getenv("DB_PASS")
	dbName   = os.Getenv("DB_NAME")
)

func initializeDatabase() {
	connString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName)

	db, err := sql.Open(dbDriver, connString)
	if err != nil {
		log.Fatalln(err)
	}

	if err := db.Ping(); err != nil {
		log.Fatalln(err)
	}

	dbConn = db
}

func saveComment(comment *CreateCommentRequest) error {
	result, err := dbConn.Exec("INSERT INTO mensagens(nome, email, comentario) VALUES(?, ?, ?)", comment.Name, comment.Email, comment.Comment)
	if err != nil {
		return err
	}

	comment.Id, err = result.LastInsertId()
	if err != nil {
		return err
	}

	return nil
}
