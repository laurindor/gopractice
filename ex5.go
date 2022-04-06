package main

import (
	"fmt"
)

type any int

var anyActual any

func main() {
	anyActual = 2022
	fmt.Println(anyActual)      //2022
	fmt.Printf("%T", anyActual) //main.any
}
