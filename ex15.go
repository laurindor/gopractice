//Ara realitza un cuadrat emprant el caràcter * un ùnic cop i que és repeteixi 10 cops d'ample i 10 cops d'alt.
package main

import (
	"fmt"
)

func main(){

	c := "*"

	for i := 0; i < 10; i++{
		for j := 0; j < 10; j++{
			fmt.Print(c, " ")
		}
		fmt.Println()
	}
}