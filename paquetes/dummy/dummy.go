package dummy

import "fmt"

func HolaMundo() string {
	return "Hola mundo :D"
}

func init() {
	fmt.Println("Se activó el paquete")
}
