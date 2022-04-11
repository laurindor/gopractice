package main

import (
	"fmt"
)

func main() {
	//A continuació realitza 2 structs anonims sobre articles de supermercats a on has de definir les dades de nom, unitats i pvp.
	article1 := struct {
		nom string
		pvp, unitats int
	}{
		nom: "ElDiari",
		pvp: 2,
		unitats: 80,
	}

	article2 := struct{
		nom string
		pvp, unitats int

	} {
		nom: "llibre"
		pvp: 10,
		unitats: 20,
	}

	fmt.Println(article1)
	fmt.Println(article2)
}
