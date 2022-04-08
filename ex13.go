//Mostra tots el nombtres del 0 al 100 de 2 en dos.Definint tres paràmetres (init, condition, post)
//pt2 que la iteracio es pari en el num 50

package main

import (
	"fmt"
)

func main(){
	surt := false
for i :=0; i <= 100 && surt == false; i+=2{
	fmt.Println(i)
	if i == 50{
		surt = true
	}
}
}


	
