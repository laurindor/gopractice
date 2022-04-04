package main

import "fmt"

func main() {
	var a int
	var b string
	var c float64
	var d bool

	fmt.Printf("var a %T = %+v\n", a, a)   //var a int =  0
	fmt.Printf("var b %T = %q\n", b, b)    //var b string = ""
	fmt.Printf("var c %T = %+v\n", c, c)   //var c float64 = 0
	fmt.Printf("var d %T = %+v\n\n", d, d) //var d bool = false
}
