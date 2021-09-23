package main

import "fmt"

type hechicero struct {
	mana   int
	nombre string
}

/*
 Se define un comportamiento es decir en este caso se va a
 escribir bytes en algun lugar ya sea un fichero una conexión TCP
 solo nos interesa lo que entra y lo que sale de la función
 esto es lo que denominamos "comportamiento"

 Lo que devuelva será (nº de bytes escritos y el posible
	error que pueda producirse) respectivamente
*/
type Escritor interface {
	Write([]byte) (int, error)
}

/*
 Ahora solo necesitamos un struct que implemente esa interfaz
 si vienes de otros lenguajes tipo java esperarías algo tipo
 implements para implementar la interfaz pero en go , eso se hace
 creando un método con la misma firma que el método de la interfaz
 que acepte este struct
*/
type EscritorConsola struct{}

/*
 Aqui definimos la implementación de la interfaz
*/
func (escritorConsola EscritorConsola) Write(bytesEscribir []byte) (int, error) {
	bytesEscritos, error := fmt.Print(string(bytesEscribir))
	return bytesEscritos, error
}

func main() {

	// De tal manera que puedo sustituir EscritorConsola{} por por ejemplo
	// EscritorArchivo{} ya que ese tambien implementaría la interfaz Escritor
	// porque cumple el contrato(método) de la misma
	var escritor Escritor = EscritorConsola{}
	escritor.Write([]byte("Hola mundo"))

}
