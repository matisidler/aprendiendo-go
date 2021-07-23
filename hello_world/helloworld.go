package main

import (
	"fmt"
	"strconv"
)

func main() {
	//Variables:
	//var Nombre_Variable Tipo_Dato
	var x, y, z int = 93, 92, 91
	var cadena string = "Hola"
	nombre := "Chau"
	nombre = "Hola"
	fmt.Println(x, y, z, cadena, nombre)

	edad := 22
	//Convertir una variable int en string
	edad_str := strconv.Itoa(edad)
	fmt.Println("Hola, mi edad es " + edad_str)
	fmt.Println(edad + 10)
	edad_int, _ := strconv.Atoi(edad_str)
	println(edad_int)
}
