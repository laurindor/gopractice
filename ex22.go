package main

import (
	"fmt"
)

func main() {
	x := make([]int, 9, 10)  //indiquem tipus, capacity y length

	x[2] = 3 //a la posicio dos em poses un 3
	x[6] = 7 //a la posicio 6 em poses un 7

	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

	x = append(x, 10) //afegim valor10, l'afegeix al final
	fmt.Println(x)
	fmt.Println(len(x)) //la length ara és 10 perquè s'ha afegit un numero mes
	fmt.Println(cap(x))

	//x[11] = 11 //Això donarà error
	x = append(x, 11)
	fmt.Println(x)
	fmt.Println(len(x)) 
	fmt.Println(cap(x)) //Això duplicarà el espai de memòria

}
