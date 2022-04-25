package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type Usuaris struct {
	ID   int    `json:"id"`
	Nom  string `json:"nom"`
	Clau string `json:"clau"`
}

func main() {
	// Conexió a la bd ciber
	db, err := sql.Open("mysql", "ciber:ciber1234@tcp(127.0.0.1:3306)/ciber")

	//Si hi ha algun error durant la conexió queda registrat
	if err != nil {
		panic(err.Error())
	}

	// deferim el tencament de la conexió, fins que el proces principal no hagi finalitzat
	defer db.Close()

	// exemple de Query per insertar un registre
	insert, err := db.Query("INSERT INTO usuaris VALUES ( 4, 'TEST', '1234' )")

	// registra els possibles errors
	if err != nil {
		panic(err.Error())
	}
	// Finalitza el Qery un cop finalitzi el procés principal
	defer insert.Close()

	//Query select *
	results, err := db.Query("SELECT id, nom, clau FROM usuaris")
	if err != nil {
		panic(err.Error()) // captura de errors
	}

	for results.Next() {
		var usuaris Usuaris
		// per cada fila, escanejem el resultat i el presentem seguint l'estructura del Struct
		err = results.Scan(&usuaris.ID, &usuaris.Nom, &usuaris.Clau)
		if err != nil {
			panic(err.Error()) // captura de errors
		}
		// i mostrem el valor del camp nom
		fmt.Println(usuaris.Nom)
	}

}
