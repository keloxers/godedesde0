package ejercicios

import (
	"strconv"
)

func Ejercicio(texto string) (int, string) {

	i, err := strconv.Atoi(texto)

	if err != nil {
		panic(err)
	}

	if i > 100 {
		return i, "es mayor a 100"
	} else {
		return i, "es mejor a 100"
	}

}
