package main

import "fmt"

func Suma(a, b int) int {
	return a + b
}

func main() {
	var x, y int
	x = 3
	y = 6
	a := Suma(x, y)
	fmt.Println("El valor de ", x, "+", y, "=", a)

}
