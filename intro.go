

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

- - - - -- - 

CICLES FOR

+ una única condició que retorna true or false, ojo generant bucles infinits

for i<10{
		
}

+ amb parámetres(exercici 13): init, condition, post:

for i:=0; i<=100;i++{
	fmt.Println(i)
}

+ com a cicle infinit (no te cap paràmetre, sempre es unfinit  s'ha de definir e
l break):

i := 0
for {

	if i<10{
		i++
	}else{
		break; 
	}
}

- - - - - 

Anidats - exercici 15


Arrays

var x[5]in  - hem d'indicar el tipus de valor

Range

recorrer arrays automàticament, fent servir el for

nums := []int{24, 62, 7, 17, 3}

for i,v := range nums{
	fmt.Println(i, "",v) 
}

Slice

Es fa servir literal compost 

//tipus{elements}//COMPOSITE LITERALS
x := []int{1,2,3,4,5}
fmt.Println(x[1:3])//[23]


Append

Es pot eliminar elements d’un Slice emprant append i slicing per a dividir.

x := []int{1, 2, 3, 4, 5}
x = append(x[:1], x[3:]...)
fmt.Println(x) // [1 4 5]


Make

(ex22) crea slices amb tres atributs: tipus, longitud (len, els elements que es rpresenten), capacitat (límit d'elements, aforo)

Slice multidimensionals

sl := []string{“Samarreta”, “Lacoste”, “Talla L”, “Blava”}
fmt.Println(sl)
pd := []string{“Polo”, “Decathlon”, “Talla XL”, “Verd”}
fmt.Println(pd)
carro := [][]string{sl, pd}
fmt.Println(carro)

Datamap

map[index]valor

m := map[string]int{
	"Josep":6,
	"Joana":8,
	"Enric":4, //l'ultim parell de claus ha de portar coma tb
}
	fmt.Println(m)
	fmt.Println(m["Josep"])






*/


