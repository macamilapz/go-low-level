package main

import (
	"fmt"
	"math/rand/v2"
	"os"
	"strconv"
)

func main() {
	var args []string = os.Args
	if len(args) != 2 {
		panic("Debes ingresar un número")
	}

	var max int = 3
	var min int = 1

	jugador, err := strconv.Atoi(args[1])
	if err != nil || (jugador < min || jugador > max) {
		panic("Debes ingresar un numero entero entre 1 y 3")
	}

	var pc int = rand.IntN(max-min) + min

	if jugador == 1 {
		fmt.Println("Elegiste ✊")
	} else if jugador == 2 {
		fmt.Println("Elegiste✋")
	} else {
		fmt.Println("Elegiste ✌")
	}

	if pc == 1 {
		fmt.Println("PC eligió ✊")
	} else if pc == 2 {
		fmt.Println("PC eligió ✋")
	} else {
		fmt.Println("PC eligió ✌")
	}

	if pc == jugador {
		fmt.Println("Empate")
	} else if jugador == 1 && pc == 3 {
		fmt.Println("Ganaste")
	} else if jugador == 3 && pc == 2 {
		fmt.Println("Ganaste")
	} else if jugador == 2 && pc == 1 {
		fmt.Println("Ganaste")
	} else {
		fmt.Println("Perdiste")
	}
}
