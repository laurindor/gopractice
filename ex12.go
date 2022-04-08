//mostra els núm parells de l'1 al 20

package main

import (
	"fmt"
)

func main() {
	i := 0

	for i < 20 {
		{   
			fmt.Println(i)
			i +=2 
		}
	}
}
