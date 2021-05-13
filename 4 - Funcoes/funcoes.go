package main

import "fmt"

func somar(num1 int8, num2 int8) int8 {
	return num1 + num2
}

func calculosMatematicos(n1, n2 int8) (int8, int8) {
	soma := n1 + n2
	sub := n1 - n2
	return soma, sub
}

func main() {
	soma := somar(10, 20)
	fmt.Println(soma)

	var f = func(txt string) string {
		return txt
	}

	resultado := f("Texto da função 1")
	fmt.Println(resultado)

	soma, subtracao := calculosMatematicos(10, 15)
	fmt.Println(soma, subtracao)

	// Quando informado _ ele entende que você não vai utilizar a váriavel
	somat, _ := calculosMatematicos(10, 15)
	fmt.Println(somat)
}
