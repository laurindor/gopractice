package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	acces   bool = false
	userIn  string
	claveIn string
	i       int
)

const (
	user  = "admin"
	clave = "123456"
)

func main() {

	for i = 3; i > 0 && acces == false; i-- {
		fmt.Println("Ciber Login. Disposes de ", i, " intents.\nIndica el teu usuari?")
		userIn = scanner()
		fmt.Println("Indica la teva clau?")
		claveIn = scanner()
		//calculem la longitud d'un string amb len()
		if len(claveIn) < 4 || len(claveIn) > 6 {
			fmt.Println("Error. Tens que indicar una clau amb una longitud d'entre 4 i 6")
			continue
		}

		//validacio de les credencials
		if user == userIn && clave == claveIn {
			fmt.Println("Acces correcte")
			break
		} else {
			fmt.Println("Dades incorrectes")
		}
	}
	if i == 0 {
		fmt.Println("Ha excedit el numero de intents")
	}
}

func scanner() string {
	reader := bufio.NewScanner(os.Stdin)
	//text, _ := reader.ReadString('\n')
	reader.Scan()
	return reader.Text()

}
