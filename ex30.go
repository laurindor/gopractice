package main

import (
	"fmt"
)

func multi(num int) int {
	return num * num
}

func cuadrat(f func(int) int, list []int) []int {
	// Make serveix per crear Slices i s’ha de indicar make(tipus[], longitud, capacitat)
	var a = make([]int, len(list), cap(list))
	for index, val := range list {
		a[index] = f(val)
	}
	return a
}

func main() {
	llistat := []int{3, 5, 8, 10}
	resultat := cuadrat(multi, llistat)
	fmt.Println(resultat)

}
