package main

import (
	"fmt"
	s "strings"
)

func main() {
	frase := "test"
	//Revisem si un  esta contingut dins d'un altre string al final
	fmt.Println(s.HasSuffix(frase, "st"))
}
