//Exercici crear conexió, configurar usuari, crear bd, crear taula i insertar registre

package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// The database is called testDb
	db, err := sql.Open("mysql", "ciber:ciber1234@tcp(127.0.0.1:3306)/ciber")

	// if there is an error opening the connection, handle it
	if err != nil {
		panic(err.Error())
	}

	fmt.Println(db)
	// defer the close till after the main function has finished
	// executing
	defer db.Close()

	// perform a db.Query insert
	insert, err := db.Query("INSERT INTO usuaris VALUES ( 2, 'TEST', '1234' )")

	// if there is an error inserting, handle it
	if err != nil {
		panic(err.Error())
	}
	// be careful deferring Queries if you are using transactions
	defer insert.Close()
}
