package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	var args []string = os.Args

	if len(args) != 2 {
		panic("use enteros")
	}

	jugador, err := strconv.Atoi(args[1])

	if err != nil || (jugador < 1 || jugador > 3) {
		panic("Debes ingresar un numero entero entre 1 y 3")
	}
	fmt.Println(jugador)
}
