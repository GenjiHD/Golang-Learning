package main

import (
	"fmt"
	"strings"
)

// Como se escribe una funcion

func main() {
	fmt.Println("Hello World")

	// Numeros
	// var entero int
	entero := 10
	decimal := 3.14
	suma := entero + int(decimal)

	// texto
	mensaje := "Hello, "
	concatenado := mensaje + "World"
	enMayusculas := strings.ToUpper(concatenado)

	// Booleanos
	esVerdadero := true

	// Arrays y Slices
	arrayFijo := [2]int{1, 2} // Tenemos que colocar una longitud fija en este caso

	// Solucion a esto
	sliceVariable := []int{4, 5, 6}
	sliceVariable = append(sliceVariable, 7)

	// Mapas
	diccionario := map[string]int{
		"clave1": 1,
		"clave2": 2,
	}

	//Structs
	// Igual que los types de TS
	type Persona struct {

		// Definir elementos
		Nombre string
		Edad   int
	}

	persona := Persona{Nombre: "Joaquin", Edad: 21}

	fmt.Println("Suma, ", suma)
	fmt.Println("Mensaje, ", enMayusculas)
	fmt.Println("Array, ", arrayFijo)
	fmt.Println("Slice, ", sliceVariable)
	fmt.Println("Mapa, ", diccionario)
	fmt.Println("Persona, ", persona)
	fmt.Println("Booleano, ", esVerdadero)
}

// Tipos de datos

// int, int64, int 32, int16, int8 | entero | controlar el size de los enteros |
// uint, uint16, uint32, uint64 | enteros sin signo | cuando no queremos negativos | 0
// float32, float64 | cuando queremos representar valores numericos reales | numeros con decimales => 32 | 64 asociado al sistema | 0
// bool true/false | flag o condicionales | false
// string | para representar texto | null
// byte == uint8 | trabajar con datos binarios
// rune == int32 | cuando queremos representar un solo caracter que ocupa mas de 1 byte
// complex64, complex128 | cuando tiene una parte real y una parte imaginaria | N + iN | 0 + 0i
