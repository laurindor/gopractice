package main

import (
	"fmt"
	s "strings"
)

func main() {
	frase := "test"
	//Revisem si un string esta contingut dins d'un altre string
	fmt.Println(s.Contains(frase, "es"))
}
