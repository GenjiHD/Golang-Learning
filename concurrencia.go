package main

import (
	"fmt"
	"sync"
	"time"
)

// go => GoRoutine
// GoRoutine => es un hilo de ejecucion ligero virtual, asigna una tarea a cada hilo para aprovechar
// todo lo que puedo del CPU, osea puedo aprovechar un servidor entero
// defer metodo, se ejecuta al final de la funcion
// Channel => comunicar cosas entre GoRoutines

// IMPORTANTE: como poner el tipo de canal, es lo que se puede hacer con el mismo
func decirHola(canal chan<- string) {
	time.Sleep(1 * time.Second)
	canal <- "Hola desde la GoRoutine"
}

func imprimirMensaje(canal <-chan string) {
	fmt.Println("Esperando Mensaje...")
	msg := <-canal
	fmt.Println(msg)
}

func main() {
	canal := make(chan string)
	go decirHola(canal)
	imprimirMensaje(canal)

	// make no solo sirve para crear algo como lo slice
	// tambien sirve para crear canales u otra cosa
	canal2 := make(chan int)
	go func() {
		for i := 0; i < 5; i++ {
			canal2 <- i
		}

		close(canal2)
	}()

	// range puede tanto leer como recibir
	for num := range canal2 {
		fmt.Println("Numero recibido: ", num)
	}

	// Ejemplo Mutex
	var contador int
	var mu sync.RWMutex // tanto escritor como lector al mismo tiempo

	// Mutex es para reservar recursos

	// Writer
	go func() {
		for i := 0; i < 5; i++ {
			mu.Lock()
			contador++
			mu.Unlock()
			time.Sleep(100 * time.Millisecond)
		}
	}()

	// Reader
	for i := 0; i < 3; i++ {
		go func(id int) {
			for j := 0; j < 5; j++ {
				mu.RLock()
				fmt.Println("Leyendo desde la GoRoutine %d: %d\n", id, contador)
				mu.RUnlock()
				time.Sleep(100 * time.Millisecond)
			}
		}(i)
	}

	time.Sleep(2 * time.Second)
}
