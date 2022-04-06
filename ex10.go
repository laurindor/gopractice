//Heu de realitzar una estructura condicional Switch que contempli els seguents case basant-se en un avatar fictici:
// 1. Tingui energia al 100% o tingui un booster, mostrant el missatge "Energia carregada"
// 2. Tingui la energia menor a 100, augmentant l'energia en 20 unitats i indicant el nivell actual d'energia.
// 3. Un default, mostrant un error.
package main

import "fmt"

func main() {
	energia := 78
	booster := false

	switch {
	case energia == 100, booster == true:
		fmt.Println("Energia carregada")

	case energia < 100:
		energia += 20 // o pots posar energia = energia + 20
		fmt.Println("Energia actual" , energia, "després de carregar-te 20 unitats")
	
	default:
		fmt.Println("Nivell d'energia incorrecte")
	}

}
