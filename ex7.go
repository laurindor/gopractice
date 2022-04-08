package main

import (
	"fmt"
)

func main() {
	edat := 17
	dia := "dissabte"

	if edat > 18 || (dia == "dijous" && edat >= 16) {
		fmt.Println("Pots entrar")
	} else if edat == 0 {
		fmt.Println("M'has de dir la teva edat")
	} else {
	fmt.Println("No pots entrar")
	}
}
