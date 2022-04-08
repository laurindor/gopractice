package main

import "fmt"

func main() {
	carro := []string{"pastanagues", "pebrots", "aigua", "formatge"}
	//Sobre el carro de la compra que et mostrem, afegeix un cogombre en la posició 1 substituint l'altre article.
	carro[1] = "cogombre"

	fmt.Println(carro)
	i := 0
	for i < len(carro){ 
		fmt.Println(carro[i])
		i++



	}
}
