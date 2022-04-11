package main

import (
	"fmt"
)


func main(){
	//slice				//tipus, len, cap
	notesalumnes := make([]int,10,100)
	fmt.Println(notesalumnes)
	fmt.Println(len(notesalumnes))
	fmt.Println(cap(notesalumnes))
}