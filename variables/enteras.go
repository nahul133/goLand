package variables

import (
	"fmt"
)

func MostrarEnteros() {
	var intcomun int
	intde32 := int32(10)
	intde64 := int64(100)

	fmt.Println("intComun =", intcomun)
	fmt.Println("int32 =", intde32)
	fmt.Println("int64 =", intde64)
}
