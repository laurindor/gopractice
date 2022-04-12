package main

import (
	"fmt"
)

type alumne struct {
	nom    string
	cognom string
}
type matriculat struct {
	alumne
	asignatura string
}

//Exemple de mètode
func (a alumne) presentarse() { //D’aquesta manera afegim aquest mètode al struct matriculat
	fmt.Println("El meu nom és ", a.nom, a.cognom, " i estic aprenent molt.")
}
func (m matriculat) presentarse() { //D’aquesta manera afegim aquest metode al struct matriculat
	fmt.Println("El meu nom és ", m.nom, m.cognom, " i estic impartint ", m.asignatura)
}

type usuari interface {
	presentarse()
}

func aula(u usuari) {
	switch u.(type) {
	case alumne:
		fmt.Println("Estic dins de l'aula", u.(alumne).nom)
	case matriculat:
		fmt.Println("Estic dins de l'aula", u.(matriculat).nom)
	}
	fmt.Println("Estic dins de l'aula", u)
}
func main() {
	al1 := matriculat{
		alumne: alumne{
			nom:    "Jon",
			cognom: "Ruiz",
		},
		asignatura: "Go",
	}

	pr1 := alumne{
		nom:    "Joan",
		cognom: "Riera",
	}

	fmt.Println(al1)  //{{Jon Ruiz} Go}
	al1.presentarse() //El meu nom és Jon Ruiz i estic aprenent Go
	pr1.presentarse() //El meu nom és  Joan Riera  i estic aprenent molt.
	aula(pr1)         //Estic dins de l'aula {Joan Riera}
	aula(al1)         //Estic dins de l'aula {{Jon Ruiz} Go}
}
