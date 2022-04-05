package main

import (
	"fmt"
)

var ciutat string = "Barcelona"

func main() {
	frase := fmt.Sprint("La ciutat de ", ciutat, " està a Catalunya")
	fmt.Println(frase) //La ciutat de Barcelona està a Catalunya

}
