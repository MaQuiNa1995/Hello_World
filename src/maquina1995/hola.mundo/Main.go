package main

import (
	"fmt"
)

func main() {

	// copia nueva
	profesiones := [...]string{"Caballero Cebolla", "Ilusionista", "Francotirador"}
	profesionesCopia := profesiones
	profesionesCopia[0] = "Mago Azul"
	fmt.Printf("Profesiones:       %v\nProfesiones Copia: %v\n", profesiones, profesionesCopia)

	// copia real apuntando a la misma dirección de memoria
	profesiones2 := [...]string{"Caballero Cebolla", "Ilusionista", "Francotirador"}
	profesionesCopia2 := &profesiones2
	profesionesCopia2[0] = "Domador"
	fmt.Printf("Profesiones:       %v\nProfesiones Copia: %v\n", profesiones2, profesionesCopia2)

}
