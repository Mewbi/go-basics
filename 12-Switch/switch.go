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
		return "Sábado-Feira"

	default:
		return "Número inválido"
	}
}

func diaDaSemana2(numero int) string {
	var diaDaSemana string
	switch {
	case numero == 1:
		diaDaSemana = "Domingo"
		fallthrough
	case numero == 2:
		diaDaSemana = "Segunda-Feira"

	case numero == 3:
		diaDaSemana = "Terça-Feira"

	case numero == 4:
		diaDaSemana = "Quarta-Feira"

	case numero == 5:
		diaDaSemana = "Quinta-Feira"

	case numero == 6:
		diaDaSemana = "Sexta-Feira"

	case numero == 7:
		diaDaSemana = "Sábado-Feira"

	default:
		diaDaSemana = "Número inválido"
	}

	return diaDaSemana
}

func main() {
	fmt.Println("Switch")

	dia := diaDaSemana(20)
	fmt.Println(dia)

	dia2 := diaDaSemana(1)
	fmt.Println(dia2)
}
