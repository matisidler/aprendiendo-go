package main

import (
	"fmt"
)

func main() {
	edades := map[string]int{
		"ana":    22,
		"juan":   11,
		"matias": 19,
	}
	delete(edades, "ana")
	edades["matias"] += 2
	fmt.Println(edades)

	for nombre, edad := range edades {
		fmt.Printf("La edad de %s es %d \n", nombre, edad)
	}
}
