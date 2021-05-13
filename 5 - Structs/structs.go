package main

import "fmt"

type usuario struct {
	nome     string
	idade    uint8
	endereco endereco
}

type endereco struct {
	logradouro string
	numero     uint8
}

func main() {
	fmt.Println("Arquivo structs.")

	var u usuario
	u.nome = "Johann"
	u.idade = 25

	fmt.Println(u)

	usuario2 := usuario{"Johann", 25, endereco{"Rua 1", 2}}
	fmt.Println(usuario2)

	usuario3 := usuario{nome: "Johann"}
	fmt.Println(usuario3)
}
