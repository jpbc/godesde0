package ejercicios

import "strconv"

func DevuelveIntString(cadena string) (int, string) {
	entero, err := strconv.Atoi(cadena)
	if err != nil {
		// ... handle error
		panic(err)
	}

	var mensaje string
	if entero > 100 {
		mensaje = "Es Mayor a 100"
	} else {
		mensaje = "Es Menor a 100"
	}
	return entero, mensaje
}
