package main

import (
	"fmt"
)

func main() {
	edat := 19
	if edat >= 18 {
		fmt.Println("Pots accedir") //Missatge per quan es compleix la condició
	} else {
		fmt.Println("No pots passar") //Missatge per quan no es compleix la condició
	}
}