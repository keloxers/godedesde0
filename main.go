package main

import (
	"fmt"

	"github.com/keloxers/godedesde0/ejercicios"
)

//"fmt"
// "runtime"
// "github.com/keloxers/godedesde0/variables"

func main() {

	// variables.MuestroEnteros()
	//variables.RestoVariables()

	// estado, texto := variables.ConviertoaTexto(1588)
	// fmt.Println((estado))
	// fmt.Println(texto)

	// if os := runtime.GOOS; os == "darwin" || os == "Linux" {
	// 	fmt.Println("No Windows, es", os)
	// } else {
	// 	fmt.Println("Windows")
	// }

	// switch os := runtime.GOOS; os {
	// case "darwin":
	// 	fmt.Println("MacOS")
	// case "Linux":
	// 	fmt.Println("Linux")
	// default:
	// 	fmt.Println("Windows")
	// }

	fmt.Println(ejercicios.Ejercicio("200"))

}
