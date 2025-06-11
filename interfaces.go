package main

import (
	"fmt"
	"math"
)

// las interfaces NO se implementan al menos manualmente
// las interfaces se cumplen

type Forma interface {
	Area() float64
}

type Circulo struct {
	Radio float64
}

func (c Circulo) Area() float64 {
	return math.Pi * c.Radio * c.Radio
}

func imprimirArea(f Forma) {
	fmt.Println("El area es: %.2f\n", f.Area())
}

func main() {
	c := Circulo{Radio: 5}
	imprimirArea(c)

	var cualquiera interface{} // typescript any = "esto es un string"
}
