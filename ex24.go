package main

import (
	"fmt"
)

type persona struct {
	nom  string
	edat int
}

func main() {
	p1 := persona{
		nom:  "Gerard",
		edat: 22,
	}

	p2 := persona{
		nom:  "Ona",
		edat: 18,
	}

	fmt.Println(p1)
	fmt.Println(p2)
	fmt.Println(p1.nom)
	fmt.Println(p2.nom)
}
