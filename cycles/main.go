package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	// Declaramos una variable de tipo *bufio.Reader
	var reader *bufio.Reader = bufio.NewReader(os.Stdin)
	// Indicamos al usuario que debe hacer
	fmt.Printf("¿Hasta donde quiere contar sumerce?: ")
	// Leemos lo que el usuario escriba hasta que le de enter
	in, _ := reader.ReadString('\n')
	// quitamos los dos ultimos caracteres para quede sin el "\n"
	in = in[:len(in)-2]

	// convertimos el input a numero
	count, err := strconv.Atoi(in)
	if err != nil {
		panic("debes escribir un numero entero positivo")
	}

	for i := 1; i <= count; i++ {
		fmt.Println(i)
	}

}
