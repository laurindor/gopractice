package main

import (
	"fmt"
)

func main() {
	//Mira d'obtenir el total de la suma dels numeros parells de la col.lecció que indiquem i només emprant una estructura for sense atributs.
	nums := []int{3, 6, 12, 7, 5, 8, 2}

	i, total := 0,0
	for{
		if i == len(nums){
			break
		}
		if nums[i] % 2 == 0{
			total += nums[i]
		} 
		i++
	}
	fmt.Println("La suma dels números parells és ", total)

}
