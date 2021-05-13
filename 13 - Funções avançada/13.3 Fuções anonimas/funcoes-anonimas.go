package main

import "fmt"

func main() {

	func(texto string) {
		fmt.Println(texto)
	}("Passando parametro")

	texto := func(texto string) string {
		return fmt.Sprintf("Recebido -> %s", texto)
	}("Passando parametro")

	fmt.Println(texto)
}
