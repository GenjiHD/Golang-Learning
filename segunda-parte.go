package main

import "fmt"

func main() {
	defer fmt.Println("FIN")
	// Condicional
	// Assertive / negative programming
	edad := 20
	if edad < 18 {
		fmt.Println("Eres menor de edad")
		return
	}

	fmt.Println("Eres mayor de edad")

	// Bucles classic
	for i := 0; i < 5; i++ {
		fmt.Printf("Iteracion: %d\n", i)
	}

	// Bucle tipo While
	// n := 0
	// for n < 3 {
	// 	fmt.Printf("Iteracion: %d\n", n)
	// 	n++
	// }

	// Bucle infinito
	n := 0
	for {
		n++

		if n == 5 {
			continue
		}

		fmt.Printf("Nn bucle infinito: %d\n", n)

		if n >= 7 {
			break
		}
	}

	// Range
	// Nos permite recorrer todo
	slice := []string{"uno", "dos", "tres"}

	for index, value := range slice {
		fmt.Printf("Indice %d, Valor: %s\n", index, value)
	}

	// Switch
	valor := 2
	switch valor {
	case 1:
		fmt.Println("Es 1")
	case 2:
		fmt.Println("Es 2")
	default:
		fmt.Println("No es 1 ni 2")
	}
}
