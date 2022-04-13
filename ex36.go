package main

import (
	"fmt"
	s "strings"
)

func main() {
	frase := "test"
	//Revisem si un string esta contingut dins d'un altre string al inici
	fmt.Println(s.HasPrefix(frase, "te"))
}
