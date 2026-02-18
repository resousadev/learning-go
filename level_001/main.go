package main

import "fmt"

func main() {

	x := 42
	y := "James Bond"
	z := true

	// %d para numeros de base 10, %s para strings e %t para booleanos
	fmt.Printf("Os valores de x, y e z são respectivamente: %d, %s e %t\n", x, y, z)

	fmt.Println("O valor de x é:", x)
	fmt.Println("O valor de y é:", y)
	fmt.Println("O valor de z é:", z)

}
