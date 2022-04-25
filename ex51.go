package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	var jsonBlob = []byte(`[
		{"Usuari": "ciber", "Clau": "123456"},
		{"Usuari": "narium", "Clau": "qwerty"}
		]`)
	type Usuari struct {
		Usuari string
		Clau   string
	}
	var usuaris []Usuari
	//Decodifiquem el json amb Unmarshal
	err := json.Unmarshal(jsonBlob, &usuaris)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Printf("%+v", usuaris)
}
