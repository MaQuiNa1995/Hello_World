package main

import (
	"fmt"
)

const (
	constanteEnumerada  = iota + 5 // valdría 5
	constanteEnumerada2            // valdría 6
)

const (
	_  = iota
	KB = 1 << (10 * iota) // Con esto conseguimos una potencia 10^1
	MB                    // Con esto conseguimos una potencia 10^2 etc
	GB
	TB
	PB
	EB
	ZB
	YB
)

func main() {
	// Aqui tenemos un decimal que representa el peso de un fichero
	fileSize := 4000000000.
	/*
		Con el Printf aplicamos cierto formato
		%.2f -> decimal de precisión 2
		concatenado con " GB"
	*/
	fmt.Printf("%.2f GB", fileSize/GB) // Imprime 3.73 GB
}
