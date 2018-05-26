package main

import "fmt"

func add(x, y float64) float64 {
	return x + y
}

func multiple(a, b string) (string, string) {
	return a, b
}

func someFunction() {
	var num1, num2 float64 = 5.6, 9.5
	word1, word2 := "Hello", "World"

	fmt.Println(add(num1, num2))
	fmt.Println(multiple(word1, word2))
}

func casting() {
	// Casting
	var a int = 62
	var b float64 = float64(a)

	x := a // x is int

	fmt.Println(a, b, x)
}

func main() {
	someFunction()
	casting()
}
