package main

import (
	"fmt"
)

func main() {
	fmt.Println("Loops")

	i := 0

	for i < 10 {
		//time.Sleep(time.Second)
		fmt.Println("Incrementando i")
		i++
	}
	fmt.Println(i)

	for j := 0; j < 10; j++ {
		fmt.Println("Incrementando j")
	}

	arrayNomes := [3]string{"Joao", "Davi", "Lucas"}

	for indice, nome := range arrayNomes {
		fmt.Println(indice, nome)
	}

	nome := "johann"
	for indice, letra := range nome {
		fmt.Println(indice, string(letra))
	}

	usuario := map[string]string{
		"nome":      "Leonardo",
		"sobrenome": "Silva",
	}

	for chave, valor := range usuario {
		fmt.Println(chave, valor)
	}
}
