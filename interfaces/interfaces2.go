package main

import "fmt"

type humano interface {
	comer()
	respirar()
	pensar()
	EsHombre() string
}

type hombre struct {
	nombre     string
	sexo       string
	respirando bool
	pensando   bool
	comiendo   bool
}

type mujer struct {
	nombre     string
	sexo       string
	respirando bool
	pensando   bool
	comiendo   bool
}

func (h *hombre) comer()           { h.comiendo = true }
func (h *hombre) respirar()        { h.respirando = true }
func (h *hombre) pensar()          { h.pensando = true }
func (h *hombre) EsHombre() string { return "Hombre" }

func (m *mujer) comer()           { m.comiendo = true }
func (m *mujer) respirar()        { m.respirando = true }
func (m *mujer) pensar()          { m.pensando = true }
func (m *mujer) EsHombre() string { return "Mujer" }

func HumanoRespirando(hu humano) {
	hu.respirar()
	fmt.Printf("Soy un/a %s y ya estoy respirando \n", hu.EsHombre())
}
func main() {
	pedro := new(hombre)
	HumanoRespirando(pedro)

	marta := new(mujer)
	HumanoRespirando(marta)
}
