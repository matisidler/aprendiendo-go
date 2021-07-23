package main

import "fmt"

type User struct {
	edad             int
	nombre, apellido string
}

func (this User) nombre_completo() string {
	return this.nombre + " " + this.apellido
}

func (this *User) cambiar_nombre(nombre_nuevo string) string {
	this.nombre = nombre_nuevo
	return this.nombre
}

func main() {
	usuario1 := new(User)
	usuario1.nombre = "Matias"
	usuario1.apellido = "Sidler"
	fmt.Println(usuario1.nombre_completo())

	usuario1.cambiar_nombre("Juan")
	fmt.Println(usuario1.nombre)
}
