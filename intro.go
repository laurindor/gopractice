

package main 

import (//És opcional però acostuma a ser necessari el ús d’aquesta sintaxis a on
	"fmt" // invoquem els diferents mòduls que intervindran en el nostre script.(console.log)
) //En el nostre cas hem d’emprar el fmt com a funció de sortida.


var ( //assignació múltiple de valors  https://go.dev/ref/spec#Operators_and_punctuation
	casa  = "unifamiliar"
	color = "verd"
)


//les constants no es poden modificar (poden ser amb tipus o sense tipus) Go assigna solet el tipus si no en té
const 	iva = 21
const	nomTenda = "Cybershop"
	type cartell string
	var tenda1 cartell


func main(){
tenda1 = nomTenda //ojo amb les variables locals i les variables globals
fmt.Println(tenda1)
}



func main() {//La funció main acompanya el package i només es defineix un cop, aglutinant la sintaxis principal del script.

	nom := "Lau"
	nom = "Marc" //reassignació de variable, pero ha de ser del mateix tipus que l'original (string aqui). Sempre es farà servir l'ultim valor assignat
	text, decimal, enter := "llibre", "0,5", "10"
	fmt.Println("Hola classe") //metode de sortida
	fmt.Println(nom)
	fmt.Println(text)
	fmt.Println(decimal)
	fmt.Println(enter)
	fmt.Println(casa, color)

}


/*emmagatzemar variables, li posem un nom i cada cop que la fem servir la cridem, variable tipus int

num := 58

num es la var
:= es el indicador
58 es la variable tipus int 

no es pot declarar una variable que no utilitzis!!!! Donarà error



- - - -

+ Variables Booleanas, same as JS, true and false

+ strings literal interpretats - els que estan dins de "" però sense salts de linia, /, ', 

+ strings literal no interpretats (low string literal) - inclouen tot el que esta dins de ""

+ valor zero - tots els tipus de variables poden tenir valor zero, ho interpretem així %T 


- - -- - - 


verbs

%d representacio d un valor
%T representació de tipus
%+v representació de valor
%q representació de string entre cometes(quote)



packages

fmt  (format)

strconv (stringconvert)




- - - - - -


Switch amb expressió - per defecte és true

Switch sense expressió


Fallthough - executa la llogica del case encara que no es compleixi l'expresssio
switch {
	case 2 > 6:
	fmt.Println("No s'acompleix")
	case 6 > 2:
		fmt.Println("Si s'acompleix")
		fallthough
	default:
		fmt.Println ("No ho mostris")
}




*/


