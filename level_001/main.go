package main

import "fmt"

var x int
var y string
var z bool

func main() {

	x = 42
	y = "Go is awesome!"
	z = true

	// Aqui estamos atribuindo todos os valores a uma string formatada usando Sprintf
	s := fmt.Sprintf("%d %s %t", x, y, z)
	fmt.Println(s)

}
