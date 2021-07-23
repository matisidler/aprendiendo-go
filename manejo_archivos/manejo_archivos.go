package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	modificarArchivo()
}

func leoArchivo() {
	archivo, err := ioutil.ReadFile("./archivo.txt")
	if err != nil {
		fmt.Println("Hubo un error")
	} else {
		fmt.Println(string(archivo))
	}
}

func leoArchivo2() {
	archivo, err := os.Open("./archivo.txt")
	if err != nil {
		fmt.Println("Hubo un error")
	} else {
		scanner := bufio.NewScanner(archivo)
		//Bufio nos permite escanear el archivo y recorrerlo registro por registro
		for scanner.Scan() {
			//Con el bucle le decimos que escanee linea por linea
			registro := scanner.Text()
			//La variable escaner ya tiene guardado el contenido del archivo, pero para poder leerlo en String
			//Usamos el metodo TEXT y lo guardamos en la variable registro.
			fmt.Printf("Linea > %s \n", registro)
		}

		archivo.Close()

	}
}

func crearArchivo() {
	archivo, err := os.Create("./nuevoArchivo.txt")
	//Create es lo mismo que un open, solo que me permite crear el archivo.
	if err != nil {
		fmt.Println("Hubo un error")
	} else {
		fmt.Fprintln(archivo, "Esta es una linea nueva")
		//el Fprintln hace lo mismo que un prinln con la diferencia de que se guarda en un archivo.
		//Como primer parametro le digo donde se guarda
		archivo.Close()
	}
}

func crearArchivo2() {
	fileName := "./nuevoArchivo.txt"
	if Agregar(fileName, "\n Esta es una segunda linea") == false {
		fmt.Println("Hubo un error")
	}
}

func modificarArchivo() {
	archivo, _ := os.OpenFile("./nuevoArchivo.txt", os.O_WRONLY|os.O_APPEND, 664)
	archivo.WriteString("\n Hola")
}

func verificarExiste() {
	_, err := os.Stat("./archivoPrueba.txt")
	if err != nil {
		fmt.Println("No existe ese archivo")
	}

}

func Agregar(archivo string, texto string) bool {
	arch, err := os.OpenFile(archivo, os.O_WRONLY|os.O_APPEND, 664)
	//Como siempre, creo las variables arch y err, pero en este caso uso el metodo os.OpenFile
	//Como primer parametro va el archivo que quiero abrir.
	//Como segundo parametro le digo el metodo, WriteOnly, ReadOnly, etc.
	//Con | le puedo agregar un segundo metodo, en este caso el O_Append. Esto significa que no me tiene
	//que borrar el contenido cada vez que le agregue texto.
	//Y por ultimo, el tipo de permiso para que el archivo sea modificado.

	if err != nil {
		fmt.Println("Hubo un error")
		return false
	}

	_, err = arch.WriteString(texto)
	if err != nil {
		fmt.Println("Hubo otro error")
		return false
	}
	return true
}
