package main

import "fmt"

type hechicero struct {
	mana   int
	nombre string
}

func main() {

	mago := hechicero{
		100,
		"MaQuiNa",
	}

	fmt.Println("Maná actual del mago: ", mago.mana)
	// Aunque el struct no tenga este metodo hemos anclado esa función a ese tipo
	mago.lanzarHechizo()
	fmt.Println("Maná actual del mago: ", mago.mana)
}

// en este caso estamos "añadiendo" una función para poder ser invocada desde hechicero
// de tal manera aunque los struct no puedan tener una función dentro como si de una clase
// y un método en java serían podemos vincular una función a cierto struct

// notese que si queremos que los cambios que hagamos a este struct se hagan
// efectivos para en este caso la función main desde donde hemos pasado el struct
// necesitamos marcar el tipo del parámetro con un *
func (mago *hechicero) lanzarHechizo() {
	mago.mana = mago.mana - 10
	fmt.Println(mago.nombre, " Lanza Bola de fuego maná restante: ", mago.mana)
}
