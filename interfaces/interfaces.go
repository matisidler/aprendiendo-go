package main

import "fmt"

type User interface {
	Permisos() int
	Nombre() string
}

type Admin struct {
	nombre  string
	permiso int
}

func (this Admin) Permisos() int {
	return this.permiso
}

func (this Admin) Nombre() string {
	return this.nombre
}

func auth(user User) string {
	if user.Permisos() > 4 {
		return user.Nombre() + " tiene permisos de Administrador"
	}
	return "No tiene permisos de Administrador"
}

func main() {
	admin := Admin{"Matias", 3}
	fmt.Println(auth(admin))
}
