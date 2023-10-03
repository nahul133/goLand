package teclado

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var numero1 int
var numero2 int
var leyenda string
var err error

func IngresoNumeros() {
	fmt.Println("Ingrese Numero 1 : ")
	scanner := bufio.NewScanner(os.Stdin)

	if scanner.Scan() {
		numero1, err = strconv.Atoi(scanner.Text())

		if err != nil {
			panic("El dato Ingresado es Incorrecto" + err.Error())
		}

	}

	fmt.Println("Ingrese Numero 2 : ")
	if scanner.Scan() {
		numero2, err = strconv.Atoi(scanner.Text())

		if err != nil {
			panic("El dato Ingresado es Incorrecto" + err.Error())
		}

	}

	fmt.Println("Ingrese Leyenda : ")
	if scanner.Scan() {
		leyenda = scanner.Text()

	}

	fmt.Println(leyenda, numero1*numero2)
}
