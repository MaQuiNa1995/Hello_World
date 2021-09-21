package main

import "fmt"

func main() {
	profesiones := map[int]string{
		675: "Soldado",
		432: "Paladín",
		321: "Dragontino",
		968: "Lancero",
	}
	_, ok := profesiones[675]
	fmt.Println(ok)

}
