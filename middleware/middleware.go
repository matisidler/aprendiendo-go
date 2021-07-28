package main

import "fmt"

var resultado int

func main() {
	fmt.Println("Inicio")
	resultado = middleWare(suma)(2, 3)
	fmt.Println(resultado)
	resultado = middleWare(resta)(5, 3)
	fmt.Println(resultado)
	resultado = middleWare(multiplicacion)(2, 3)
	fmt.Println(resultado)

}

func suma(a, b int) int {
	return a + b
}
func resta(a, b int) int {
	return a - b
}
func multiplicacion(a, b int) int {
	return a * b
}

func middleWare(f func(int, int) int) func(int, int) int {
	return func(a, b int) int {
		fmt.Println("Inicio de operacion")
		return f(a, b)
	}
}
