package main

import ("fmt"
)

func main() {
	var a int //numeric sencer
	var b string
	var c float64
	var d bool
				//aqui fem servir l'output printformat i no printline(Println) i posem dues vegades la variable perquè estem fent servir dos verbs
	fmt.Printf("var a %T = %+v\n", a, a)   //var a int =  0
	fmt.Printf("var b %T = %q\n", b, b)    //var b string = ""
	fmt.Printf("var c %T = %+v\n", c, c)   //var c float64 = 0
	fmt.Printf("var d %T = %+v\n\n", d, d) //var d bool = false
}
