package main

import (
	"fmt"
)

func main() {

	x := []int{1,2,3,4,5}

	b := x[:1]

	fmt.Println(b)

	c := x[3:]

	fmt.Println(c)

	//x = append(b, c[0], c[1] )
	x = append(b, c... )
	fmt.Println(x)

}