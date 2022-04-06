package main

import (
	"fmt"
)

func main() {
	a := 4
	fmt.Printf("%d\t\t%b\n", a, a) //4    100 (\n salt de linia)
	// \t imprimirá una tabulació i %b mostrara en binari
	b := a << 1
	fmt.Printf("%d binari: %b\n", b, b)
	//Hem d'imprimir un binari amb una posició més cap a la dreta
}
/* Birshifting és quan es desplacen
nombres binaris tant cap a la dreta o a l’esquerra un nombre determinat de
cops. */