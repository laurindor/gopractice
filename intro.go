

package main 

import (//És opcional però acostuma a ser necessari el ús d’aquesta sintaxis a on
	"fmt" // invoquem els diferents mòduls que intervindran en el nostre script.
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
*/


