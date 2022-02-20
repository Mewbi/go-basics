package main

import "fmt"

var n int

func main() {
	fmt.Println("Executando a função main")
	fmt.Println(n)
}

func init() {
	fmt.Println("Executando a função init")
	n = 10
}
