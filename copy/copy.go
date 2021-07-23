package main

import "fmt"

func main() {
	slice := []int{1, 2, 3, 4}
	copia := make([]int, len(slice))
	copy(copia, slice)
	fmt.Println(copia)
}
