package main

import (
	"fmt"
)

func main() {
	nota := 8

	if nota < 5 {
		fmt.Println("Suspes")
	} else if nota < 6 {
		fmt.Println("Aprobat")
	} else if nota < 9 {
		fmt.Println("Notable")
	} else {
		fmt.Println("Excelent")
	}
}
