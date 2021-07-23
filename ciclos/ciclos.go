package main

import "fmt"

func main() {
	i := 0
	for {
		fmt.Printf("Hola mundo %d\n", i)
		i++
		if i > 9 {
			break
		}
	}
}
