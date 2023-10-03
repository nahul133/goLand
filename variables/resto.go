package variables

import (
	"fmt"
	"time"
)

func MostrarRestoDeVariables() {
	Nombre := "Pedro"
	Estado := true
	Sueldo := 15755.50
	Fecha := time.Now()

	fmt.Println("String = ", Nombre)
	fmt.Println("Bolean = ", Estado)
	fmt.Println("Float = ", Sueldo)
	fmt.Println("Tipo Fecha = ", Fecha)
}
