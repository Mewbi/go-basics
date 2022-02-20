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
	fmt.Println("Arquivo Structs")

	var u usuario
	u.nome = "Davi"
	u.idade = 21
	fmt.Println(u)

	endereceExemplo := endereco{"rua nada a ver", 1}

	u2 := usuario{"Jorge", 33, endereceExemplo}
	fmt.Println(u2)

	u3 := usuario{nome: "Marcos"}
	fmt.Println(u3)
}
