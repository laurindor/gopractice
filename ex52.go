package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	s := `123456` //contrasenya de prova
	//Com podem veure el primer parametre ha de ser un slice de bytes, en el que transformem el string i el segon el cost
	bs, err := bcrypt.GenerateFromPassword([]byte(s), 4)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(s)
	fmt.Println(bs)
}
