package main

import "fmt"

func diaDaSemana(numero int) string {
	switch numero {
	case 1:
		return "Domingo"
	case 2:
		return "Segunda-Feira"
	case 3:
		return "Terça-Feira"
	case 4:
		return "Quarta-Feira"
	case 5:
		return "Quinta-Feira"
	case 6:
		return "Sexta-Feira"
	case 7:
		return "Sabado"
	default:
		return "Número Inválido"
	}
}

func diaDaSemana2(numero int) string {

	var dia string

	switch {
	case numero == 1:
		dia = "Domingo"
		fallthrough
	case numero == 2:
		dia = "Segunda-Feira"
	case numero == 3:
		dia = "Terça-Feira"
	case numero == 4:
		dia = "Quarta-Feira"
	case numero == 5:
		dia = "Quinta-Feira"
	case numero == 6:
		dia = "Sexta-Feira"
	case numero == 7:
		dia = "Sabado"
	default:
		dia = "Número Inválido"
	}
	return dia
}

func main() {
	fmt.Println("Switch")

	dia := diaDaSemana(32)
	fmt.Println(dia)

	dia2 := diaDaSemana2(1)
	fmt.Println(dia2)
}
