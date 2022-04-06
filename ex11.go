//Heu de realitzar una estructura condicional Switch que contempli els seguents case basant-se en un ascensor d'un centre comercial:
// 1. Planta primera i tercera, mostrara el missatge "Articles Llar".
// 2. Planta segona, mostrara el missatge "Moda infantil".
// 3. Planta cuarta, mostrara el missatge "Joguines"
// 4. Un default, mostrant un error.
package main

import "fmt"

func main() {

	planta := 4

	switch planta {

	case 1, 3:
		fmt.Println("Articles Llar")

	case 2:
		fmt.Println("Moda Infantil")

	case 4:
		fmt.Println("Joguines")

	default:
		fmt.Println("Error, a quina planta vas?")
	}

}
