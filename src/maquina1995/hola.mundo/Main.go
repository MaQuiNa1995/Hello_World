package main

import (
	"fmt"
	"reflect"
)

type Magia struct {
	mana   int
	nombre string `required,max:"5"`
}

type MagiaNegra struct {
	Magia
	danno int
}

func main() {
	magiaNegra := MagiaNegra{}
	magiaNegra.mana = 10
	magiaNegra.danno = 500
	// {{10} 500}
	fmt.Println(magiaNegra)
	// 10
	fmt.Println(magiaNegra.mana)
	// 500
	fmt.Println(magiaNegra.danno)

	animalReflection := reflect.TypeOf(Magia{})
	campo, _ := animalReflection.FieldByName("nombre")
	fmt.Print(campo.Tag)
}
