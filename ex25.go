package main

import (
	"fmt"
)

type persona struct {
	nom  string
	edat int
}

type ingenier struct {
	persona
	realitzarPlanol bool
}

func main() {
	ing1 := ingenier{
		persona: persona{
			nom:  "Ruben",
			edat: 26,
		},
		realitzarPlanol: true,
	}

	fmt.Println(ing1)
	fmt.Println(ing1.nom, ing1.edat, ing1.realitzarPlanol)
}
