package ejercicios

import "strconv"

func DevuelveIntString(cadena string) (int, string) {
	entero, _ := strconv.Atoi(cadena)

	var mensaje string
	if entero > 100 {
		mensaje = "Es Mayor a 100"
	} else {
		mensaje = "Es Menor a 100"
	}
	return entero, mensaje
}
