package main

import (
	"fmt"
	s "strings"
)

func main() {
	array := []string{"a", "b"}
	//Amb el mètode Join, podrem unir diferents elements d’una estructura mitjançant un nexe definit
	fmt.Println(s.Join(array, "-")) //a-b
}
