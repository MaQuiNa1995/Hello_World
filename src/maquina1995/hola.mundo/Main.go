package main

import "fmt"

func main() {
	division, error := dividir(5.0, 1.0)

	// Si error no es nulo quiere decir que hubo excepciones
	if error != nil {
		// imprimimos por consola el error
		fmt.Print("Hubo un error mas info: ", error)
		// paramos la ejecución del método
		return
	}
	// Imprimimos el valor de la división
	fmt.Print("El valor de la división es:", division)
}

// En una función podemos retornar mas de un valor
// normalmente el 1º es el resultado de la función
// y los segundos y sucesivos son errores por ejemplo
func dividir(dividendo, divisor float64) (float64, error) {

	if divisor == 0.0 {
		// En vez de controlar el flujo con panic
		// usamos la String creada a partir de fmt.Errorf
		return 0.0, fmt.Errorf("No se puede dividir entre 0")
	}

	// retornamos el valor de la división y nil
	// que indica que no hubo error
	return dividendo / divisor, nil
}
