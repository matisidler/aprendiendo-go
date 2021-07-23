package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Hola mundo")
	//PrintF me sirve para concatenar una variable a un string asi:
	// edad := 22
	//fmt.Printf("mi edad es %d", edad)
	// fmt.Println("Coloca tu edad: ")
	// fmt.Scanf("%d", &edad)
	// fmt.Printf("Mi edad es %d\n", edad)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Ingresa tu nombre: ")
	text, _ := reader.ReadString('\n')

	fmt.Printf("Hola, tu nombre es %s", text)

}
