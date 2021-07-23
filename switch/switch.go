package main

import (
	"fmt"
)

func main() {
	edad := 3

	switch edad {
	case 0:
		fmt.Println("Edad es 0")
	case 1:
		fmt.Println("Edad es 1")
	default:
		fmt.Println("No es 0 ni 1")
	}
}
