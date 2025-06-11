package domain

import (
	"fmt"
	"learning-go/domain/entities"
)

func App() {
	persona := entities.Persona{Nombre: "Joaquin", Apellido: "Castro", Edad: 21}

	fmt.Println(persona.Saludar())
	persona.CumplirAnios()

	fmt.Println(persona)
}
