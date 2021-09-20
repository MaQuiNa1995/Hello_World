package main

import "fmt"

func main() {
	slice := []string{"Mago del Tiempo", "Monje", "Paladín", "Elementalista", "Esgrimista"}
	// Crearíamos un nuevo slice omitiendo el último elemento del original
	// slice[:2] -> crea un array con: Mago del Tiempo Monje
	// slice[4:]... -> añade a ese array creado los elementos resultantes de descartar 4 elementos del slice original
	slice2 := append(slice[:2], slice[4:]...)
	// imprime: [Mago del Tiempo Monje Esgrimista]
	// se han eliminado los elementos centrales: Paladín y Elementalista
	fmt.Print(slice2)
}
