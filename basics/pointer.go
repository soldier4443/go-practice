package main

import "fmt"

// & - address
// * - value of the address ?

func main() {
	pointer()
}

func pointer() {
	x := 15
	a := &x

	// This is memory address
	fmt.Println(a)

	fmt.Println(x)
	*a = 5
	fmt.Println(x)

	fmt.Println(*a, "=", x)
}
