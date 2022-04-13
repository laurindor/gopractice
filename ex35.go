package main

import (
	"fmt"
	s "strings"
)

func main() {
	frase := "test"
	//Revisem si un string esta contingut dins d'un altre string i el numero de cops
	fmt.Println(s.Count(frase, "t"))
}
