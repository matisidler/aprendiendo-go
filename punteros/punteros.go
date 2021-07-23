package main

import "fmt"

func main() {
	var x, y *int
	entero := 5

	x = &entero //Lo que hace el & es decirle que me devuelva la dirección de memoria, NO el valor de la variable
	y = &entero

	*x = 6
	fmt.Println(*x, *y)
}
