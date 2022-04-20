package main

import (
	"fmt"
	s "strings"
)

func main() {
	frase := "a-b-c-d-e-f-g"
	//Podrem fraccionar una cadena aprofitant certs caràcters que podem emprar com a referents com a nexe.
	fmt.Println(s.Split(frase, "-")) //[a b c d e f g]
}
