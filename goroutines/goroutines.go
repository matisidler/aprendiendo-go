package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	go mi_nombre_lento("Matias")
	fmt.Println("Esperando a que termine")
	wait := ""
	fmt.Scanln(&wait)
}

func mi_nombre_lento(name string) {
	letras := strings.Split(name, "")
	for _, letra := range letras {
		time.Sleep(1000 * time.Millisecond)
		fmt.Println(letra)
	}
}
