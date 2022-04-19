package main

import (
	"fmt"
	s "strings"
)

func main() {
	frase := "test"
	//Revisem si un caràcter esta contingut dins d'un altre string i ens retorna la posició
	fmt.Println(s.Index(frase, "e")) //1
}
