package main

import (
	"fmt"
)

var edat int
var ciutat string = "Barcelona"

func main() {
	fmt.Println(edat)                                 //0
	fmt.Println(ciutat)                               //Barcelona
	fmt.Printf("El valor de edat és: %v\n", edat)     //El valor de edat és: 0
	fmt.Printf("El valor de ciutat és: %v\n", ciutat) //El valor de ciutat és: Barcelona
}
