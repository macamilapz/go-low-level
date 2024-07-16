package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	fmt.Printf(
		`seleccione una opcion:
		0. Imprimir los numeros de 1 hasta n
		1. Imprimir los numeros pares hasta n
		2. Imprimir los multiplos de 3 hasta n
		3. Imprimir los multiplos de 5 hasta n
		4. ¿ n es par ?
		5. ¿ n es multiplo de tres?
		6. ¿ n es multiplo de cinco
		7. fizzbuzz
		 `)

	// 1. Declaramos una variable de tipo *bufio.Reader
	var reader *bufio.Reader = bufio.NewReader(os.Stdin)
	// 2. Leemos lo que el usuario escriba hasta que le de enter
	in, _ := reader.ReadString('\n')
	// 3. quitamos los dos ultimos caracteres para quede sin el "\n"
	in = in[:len(in)-2]

	// 4. La opcion que eligio el usuario la convertimos a enteros
	option, err := strconv.Atoi(in)
	// 5. Se revisa si el rango de opciones es posibles de no ser asi se crea una ventana de panico avisando
	if err != nil || option < 0 || option > 9 {
		panic("Debes seleccionar una opcion valida")
	}
	// 6. cuando la opcion elegida es valida y se elige 0 preguntamos hasta donde quiere contar el usuario
	if option == 0 {
		fmt.Printf("Hasta donde quiere contar usted: ")

		//7. Repetimos linea 2 y 3 dado que necesitamos leer hasta donde quiere contar el usuario
		in, _ := reader.ReadString('\n')
		in = in[:len(in)-2]

		// 8. convertimos el input a numero si no cumple salta un panic
		count, err := strconv.Atoi(in)
		if err != nil {
			panic("debes escribir un numero entero positivo")
		}
		// 9. iteramos hasta donde eligio el usuario (i)
		for i := 1; i <= count; i++ {
			fmt.Println(i)
		}
	} else if option == 1 {
		fmt.Printf("Hasta donde quiere contar usted: ")

		in, _ := reader.ReadString('\n')
		in = in[:len(in)-2]

		count, err := strconv.Atoi(in)
		if err != nil {
			panic("debes escribir un numero entero positivo")
		}
		for i := 1; i <= count; i++ {
			if i%2 == 0 {
				fmt.Println(i)
			}
		}
	} else if option == 2 {
		fmt.Printf("Hasta donde quiere contar usted: ")

		in, _ := reader.ReadString('\n')
		in = in[:len(in)-2]

		count, err := strconv.Atoi(in)
		if err != nil {
			panic("debes escribir un numero entero positivo")
		}
		for i := 1; i <= count; i++ {
			if i%3 == 0 {
				fmt.Println(i)
			}
		}
	} else if option == 3 {
		fmt.Printf("Hasta donde quiere contar usted: ")

		in, _ := reader.ReadString('\n')
		in = in[:len(in)-2]

		count, err := strconv.Atoi(in)
		if err != nil {
			panic("debes escribir un numero entero positivo")
		}
		for i := 1; i <= count; i++ {
			if i%5 == 0 {
				fmt.Println(i)
			}
		}
	} else if option == 4 {
		fmt.Printf("Digite un numero para saber si es par: ")

		in, _ := reader.ReadString('\n')
		in = in[:len(in)-2]

		count, err := strconv.Atoi(in)
		if err != nil {
			panic("debes escribir un numero entero positivo")
		}
		if count%2 == 0 {
			fmt.Println("Tu numero es par")
		} else if count%2 != 0 {
			fmt.Println("Tu numero no es par")
		}
	} else if option == 5 {
		fmt.Printf("Digite un numero para saber si es multiplo de tres: ")

		in, _ := reader.ReadString('\n')
		in = in[:len(in)-2]

		count, err := strconv.Atoi(in)
		if err != nil {
			panic("debes escribir un numero entero positivo")
		}
		if count%3 == 0 {
			fmt.Println("Tu numero es multiplo de 3")
		} else if count%3 != 0 {
			fmt.Println("Tu numero no es multiplo de 3")
		}
	} else if option == 6 {
		fmt.Printf("Digite un numero para saber si es multiplo de 5: ")

		in, _ := reader.ReadString('\n')
		in = in[:len(in)-2]

		count, err := strconv.Atoi(in)
		if err != nil {
			panic("debes escribir un numero entero positivo")
		}
		if count%5 == 0 {
			fmt.Println("Tu numero es multiplo de 5")
		} else {
			fmt.Println("Tu numero no es multiplo de 5")
		}
	} else if option == 7 {
		fmt.Printf("fizz para multiplos de 3 y buzz para multiplos de 5: ")

		in, _ := reader.ReadString('\n')
		in = in[:len(in)-2]

		count, err := strconv.Atoi(in)
		if err != nil {
			panic("debes escribir un numero entero positivo")
		}
		for i := 1; i <= count; i++ {
			{
				if count%3 == 0 && count%5 == 0 {
					fmt.Println("FizzBuzz")
				} else if count%3 == 0 {
					fmt.Println("Fizz")
				} else if count%5 == 0 {
					fmt.Println("Buzz")
				}
			}

		}
	}
}
