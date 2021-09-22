package main

import (
	"fmt"
)

func main() {

	// numero1, numero2 := 1, 0
	// division := numero1 / numero2
	// fmt.Print(division)

	var numero1 int = 1
	var numero2 *int = &numero1
	fmt.Println(numero1, numero2) // 1 0xc0000aa058

	numero1 = 8

	// Al usar el asterisco estamos reemplazando el
	// valor de la memoria donde está almacenada por su valor real
	fmt.Println(numero1, *numero2) // 8 8

}
