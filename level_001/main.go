package main

import "fmt"

type meu_tipo int

var x meu_tipo

func main() {

	fmt.Println(x)

	fmt.Printf("O tipo de x é: %T\n", x)

	x = 42

	fmt.Printf("O valor de x é: %v\n", x)

}
