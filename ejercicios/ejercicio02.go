package ejercicios

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func LeyendoTeclado() int {
	var captura int
	var err error
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Printf("Ingrese un numero : ")
		if scanner.Scan() {
			captura, err = strconv.Atoi(scanner.Text())
			if err != nil {
				fmt.Println("Error esta digitando un caracter no valido")
				continue
			} else {
				break
			}
		}
	}
	return captura
}
func TabladeMultiplicar(numero int) {
	fmt.Println("Tabla de Multiplicar del ", numero)
	fmt.Println("==================================")
	for i := 1; i <= 10; i++ {
		if i == 10 {
			fmt.Println(i, "x ", numero, " = ", i*numero)
		} else {
			fmt.Println(i, " x ", numero, " = ", i*numero)
		}
	}
}

/* Version 2 con una sola funcion
func LeyendoTeclado() {

	var err error
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Printf("Ingrese un numero : ")
		if scanner.Scan() {
			captura, err = strconv.Atoi(scanner.Text())
			if err != nil {
				fmt.Println("Error esta digitando un caracter no valido")
				continue
			} else {
				break
			}
		}
	}
}
func TabladeMultiplicar() {
	fmt.Println("Tabla de Multiplicar del ", captura)
	fmt.Println("==================================")
	for i := 1; i <= 10; i++ {
		if i == 10 {
			fmt.Println(i, "x ", captura, " = ", i*captura)
		} else {
			fmt.Println(i, " x ", captura, " = ", i*captura)
		}
	}
}
func Ejecutar() {
	LeyendoTeclado()
	TabladeMultiplicar()
}
*/
