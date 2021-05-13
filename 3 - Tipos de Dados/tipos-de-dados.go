package main

import (
	"errors"
	"fmt"
)

func main() {
	var numero int16 = 100
	fmt.Println(numero)

	var numero2 int64 = -1000000000000000000
	fmt.Println(numero2)

	var numero3 uint32 = 10000
	fmt.Println(numero3)

	var numero4 rune = 123456
	fmt.Println(numero4)

	var numero5 byte = 123
	fmt.Println(numero5)

	var numeroReal1 float32 = 1234.45
	fmt.Println(numeroReal1)

	var numeroReal2 float64 = 1241635454.58
	fmt.Println(numeroReal2)

	numeroReal3 := 1234.88
	fmt.Println(numeroReal3)

	// String

	var str string = "Texto"
	fmt.Println(str)

	str2 := "Texto 2"
	fmt.Println(str2)

	// mais perto do char
	char := 'B'
	fmt.Println(char)

	var text int
	fmt.Print(text)

	// booleano
	var verdade bool = true
	fmt.Println(verdade)

	var mentira bool = false
	fmt.Println(mentira)

	var erro error
	fmt.Println(erro)

	var erro1 error = errors.New("Erro interno")
	fmt.Println(erro1)
}
