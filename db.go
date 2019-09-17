package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

var (
	db *sql.DB
)

func initDB() {
	var (
		host, user, password, dbname string
		port                         int
		err                          error
	)

	// TODO
	// set host, user, password db
	// also dbname

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}
}

// modify this func
func insertDB() error {
	query := `INSERT INTO table_1 ("name", "status") VALUES $1, $2`
	_, err := db.Exec(query)

	return err
}

// modify this func
func getDB() error {
	query := `SELECT "id", "name", "status" FROM table_1 WHERE "id"=$1`
	_, err := db.Query(query)

	return err
}
