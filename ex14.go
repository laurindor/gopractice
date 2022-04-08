//Realitza una estructura for que permeti recorrer tot l'array i finalitzi quan ja no hi hagin més elements per recorrer.

package main

import (
	"fmt"
)

func main(){
	alumnes := []string{"Anna", "Pep","Iria","Oscar"}

	i := 0
	for {
		switch{
			case i < len(alumnes):     //len de lenght
			fmt.Println(alumnes[i])
		
	default:
		break;
	}
	i++
	}
	
}