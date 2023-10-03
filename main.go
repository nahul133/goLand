package main

import (
	"fmt"

	"github.com/nahul133/goLand/funciones"
)

func main() {
	estado, texto := funciones.ConvertiraTexto(10)

	fmt.Println(estado)
	fmt.Println(texto)
}
