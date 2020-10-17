package main

import (
	"fmt"

	"github.com/jorgechavezrnd/mycalculator"
)

func main() {
	entrada := mycalculator.LeerEntrada()
	operador := mycalculator.LeerEntrada()
	c := mycalculator.Calc{}

	fmt.Println(c.Operate(entrada, operador))
}
