package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	type ColorGroup struct {
		ID      int
		Name    string
		Idiomas []string
	}
	group := ColorGroup{
		ID:      1,
		Name:    "Pepe",
		Idiomas: []string{"PHP", "Java", "Python", "Go"},
	}
	//Formatejar a Json el element creat a partir del struct
	b, err := json.Marshal(group)
	if err != nil {
		fmt.Println("error:", err)
	}
	os.Stdout.Write(b)
}
