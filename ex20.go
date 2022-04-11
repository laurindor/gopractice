package main

import (
	"fmt"
)

func main() {
	carro := []string{"pastanagues", "pebrots", "aigua", "formatge"}

	carro[1] = "cogombre"

	carro = append(carro, "Oli d'oliva")

	fmt.Println(carro)
}
