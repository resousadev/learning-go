package main

import "fmt"

type meu_tipo int

var x int
var y meu_tipo

func main() {

	fmt.Println(x)

	fmt.Printf("O tipo de x é: %T\n", x)

	x = 42

	fmt.Printf("O valor de x é: %v\n", x)

	y = meu_tipo(x)

	fmt.Printf("O valor de y é: %v\n", y)

	fmt.Printf("O tipo de y é: %T\n", y)

}
