package main

import "fmt"

func main() {
	recursion(2)
}

func recursion(numero int) {
	if numero > 100000000 {
		return
	}
	fmt.Println(numero)
	recursion(numero * 2)
}
