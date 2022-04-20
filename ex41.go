package main

import (
	"fmt"
	s "strings"
)

func main() {
	frase := "testejar"
	//Podrem substituir els caràcters indicats amb els paràmetres per uns altres dins d’un string.
	fmt.Println(s.Replace(frase, "e", "0", -1)) //t0st0jar
}
