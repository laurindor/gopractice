package main

import (
	"fmt"
)

var num int //Variable amb Scope a nivell de paquet
func main() {
	a := aug()
	fmt.Println(a()) //1
	fmt.Println(a()) //2
}
func aug() func() int {
	var x int
	return func() int {
		x++
		return x
	}
}

/*Al realitzar una funció anònima dins de l’altre funcio aug(),
s’està fent ús d’un espai de memòria reservat per la variable x
i que es compartirà, permetent fer servir el valor acumulat per ser autoincrementat.*/
