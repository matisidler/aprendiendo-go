package main

import "fmt"

type Human struct {
	name string
}

func (this Human) hablar() string {
	return "Bla bla bla"
}

type Animal struct {
	name string
}

type Tutor struct {
	Human
	Animal
}

func main() {
	tutor1 := Tutor{Human{"Matias"}, Animal{"Roma"}}
	fmt.Println(tutor1.Human.hablar())
}
