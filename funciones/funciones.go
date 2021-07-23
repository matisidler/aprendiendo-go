package main

import (
	"fmt"
)

func tabla(numero int) func() int {
	valor := numero
	multiplicador := 0
	return func() int {
		multiplicador++
		return valor * multiplicador
	}
}

func main() {
	MiTabla := tabla(2)

	for i := 0; i < 10; i++ {
		fmt.Println(MiTabla())
	}
}
