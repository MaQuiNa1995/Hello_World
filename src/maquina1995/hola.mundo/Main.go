package main

import "fmt"

func main() {
	saludo := "Hola"
	// nombre := "MaQuiNa"
	// nombre2 := "MaQuina2"

	saludarPorConsola(&saludo)
}

func saludarPorConsola(saludo *string, nombres ...string) {
	for _, nombre := range nombres {
		fmt.Println(*saludo, nombre)
	}
}
