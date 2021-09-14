// Declaracion del paquete
package main

// Zona de imports
import (
	"fmt"
)

// Las declaraciones a nivel de paquete no permiten el operador :=

var numero4 int
var numero5 int = 5

/*
	Adicionalmente se permite la aglutinación de distintas variables
	(que por buena practica deberían de guardar relacion entre ellas)
	con el siguiente bloque
*/
var (
	numero6 int
	numero7 int = 7
	// podemos tener el mismo nombre de variable pero en distintos "scopes"
	numero int = 1
)

/* función main de la aplicación */
func main() {

	// imprimir por pantalla
	fmt.Println("Hola Mundo")

	//Diferentes formas de declaración de variables
	//var numero float32 = 1
	//numero2 := 2
	//var numero3 int
	//numero3 = 3

	var numero float32 = 1
	/*
		El problema de esta declaracion es que nos inicializa la variable
		con float64 no hay ninguna forma de crear un float32 por esta vía
		ese sería el unico problema de esta forma
	*/
	numero2 := 2.

	// podemos aplicar cierto formato a la salida por consola
	// %v -> value de la variable
	// %T -> Tipo de la variable

	/*
		En este caso saldría 2 veces ya que tenemos 2 variables
		que se llaman numero pero con distinto scope
	*/
	fmt.Printf("%v, %T ", numero, numero)
	fmt.Printf("%v, %T", numero2, numero2)

}
