package main

import (
	"fmt"
	"time"
)

func main() {
	canal1 := make(chan time.Duration)
	go bucle(canal1)
	fmt.Println("Llegué hasta acá")

	msg := <-canal1
	fmt.Println(msg)
}

func bucle(canal1 chan time.Duration) {
	inicio := time.Now()
	for i := 0; i < 10000000000; i++ {

	}

	final := time.Now()
	canal1 <- final.Sub(inicio)
	//sub es la función que me retorna la duración
}
