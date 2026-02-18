package main

import "fmt"

var x int
var y string
var z bool

func main() {

	// Aqui será exibido o valor zero de cada variável, pois elas não foram inicializadas
	fmt.Printf("O valor de x é: %#v\n", x)
	fmt.Printf("O valor de y é: %#v\n", y)
	fmt.Printf("O valor de z é: %#v\n", z)

}
