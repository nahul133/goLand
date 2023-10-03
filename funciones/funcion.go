package funciones

import (
	"strconv"
)

func ConvertiraTexto(numero int) (bool, string) {
	texto := strconv.Itoa(numero)

	return true, texto
}
