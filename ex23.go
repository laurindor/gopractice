package main

import ("fmt"
)

func main() {
	//Realitza un map de tres alumnes i tres notes.
	notes := map[string]int{
		"Lau":7,
		"Ari":8,
		"Mire":9,
	}

fmt.Println(notes)
fmt.Println(notes["Ari"])

total := 0
for _, v := range notes{
	total += v
}
	fmt.Println("El promig és", total)

}

