package main

import "fmt"

// * -> puntero
// & => referencia | direccion de memoria

var a = 1

// NO, en verdad creamos una referencia de nombre "a" que apunta a un espacio
// de memoria que contiene en su interior el valor 1
// COPIA LOS ARGUMENTOS

func incrementar(numero *int) {
	// NUNCA modificar el arguento directamente
	// crea una copia o algo similar
	*numero++
}

// Queremos modificar el elemento original (no recomendado, es mala practica)

func main() {
	valor := 10
	fmt.Println("Valor antes de incrementar:", valor)

	incrementar(&valor)

	fmt.Println("Valor despues de incrementar:", valor)

	// new()
	puntero := new(int) //puntero int inicializado en 0
	*puntero = 20
	fmt.Println("valor inicial con new:", *puntero)
}
