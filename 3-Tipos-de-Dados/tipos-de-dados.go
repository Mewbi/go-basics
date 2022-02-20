package main

import (
	"errors"
	"fmt"
)

func main() {
	// Se não especificar o tipo do int, ele vai usar com base na arquitetura do pc, no meu caso int64
	var numero int16 = 10000
	fmt.Println(numero)

	var numero2 uint32 = 1000 // Inteiro sem sinal
	fmt.Println(numero2)

	// alias
	// int32 = rune
	var numero3 rune = 12456
	fmt.Println(numero3)

	// byte = uint8
	var numero4 byte = 123
	fmt.Println(numero4)

	var numeroReal1 float32 = 123.45
	fmt.Println(numeroReal1)

	var numeroReal2 float64 = 123123131.45
	fmt.Println(numeroReal2)

	var str string = "Texto" // não existe tipo char
	fmt.Println(str)

	char := 'B' // Se declaramos em aspas simples ele pega o valor da tabela ascii
	fmt.Println(char)

	var texto string // Todo tipo tem seu valor nulo
	fmt.Println(texto)

	var booleano1 bool = true
	fmt.Println(booleano1)

	var erro error = errors.New("Erro interno")
	fmt.Println(erro)
}
