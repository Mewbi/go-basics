package main

import (
	"fmt"
	"time"
)

func main() {
	i := 0

	for i < 10 {
		i++
		fmt.Println("Incrementando i", i)
		time.Sleep(time.Millisecond)
	}
	fmt.Println(i)

	fmt.Println("-----------")
	for j := 0; j < 10; j += 2 {
		fmt.Println("Incrementando j", j)
		time.Sleep(time.Millisecond)
	}

	nomes := [3]string{"Joao", "Davi", "Lucas"}

	fmt.Println("-----------")
	for _, nome := range nomes {
		fmt.Println(nome)
	}

	fmt.Println("-----------")
	for indice, letra := range "PALAVRA" {
		fmt.Println(indice, string(letra))
	}

	usuario := map[string]string{
		"nome":      "Leonardo",
		"sobrenome": "Silva",
	}

	fmt.Println("-----------")
	for chave, valor := range usuario {
		fmt.Println(chave, valor)
	}

	fmt.Println("-----------")
	for {
		fmt.Println("Executando infinitamente...")
		time.Sleep(time.Second)
	}
}
