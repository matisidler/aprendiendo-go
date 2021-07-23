package main

import (
	"log"
)

func main() {
	paniqueando()
}

func paniqueando() {
	defer func() {
		reco := recover()
		if reco != nil {
			log.Fatalf("ocurrió un error que generó panic \n %v", reco)
		}
	}()
	a := 1
	if a == 1 {
		panic("se encontró el valor de 1")
	}
}
