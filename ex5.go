package main

import (
	"fmt"
)

type any int

var anyActual any

func main() {
	anyActual = 2021
	fmt.Println(anyActual)      //2021
	fmt.Printf("%T", anyActual) //main.any
}
