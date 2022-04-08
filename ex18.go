package main

import (
	"fmt"
)
func main() {
	carro := []string{"pastanagues", "pebrots", "aigua", "formatge"}
	//Mira de recorrer i mostrar les dades através d'un for amb range.

	for a, b:= range carro{ //si no vull fer servir una variable la puc fer serir per un guió baix "_"
	fmt.Println(a, "",b)
	}
}