package main

import "fmt"

func main() {
	// Declaring Variables
	var i int
	i = 42
	fmt.Println(i)

	var f float32 = 43.2
	fmt.Println(f)

	// Implicid
	d := "Pesho"
	fmt.Println(d)

	bo := true
	fmt.Println(bo)

	c := complex(3, 4)
	fmt.Println(c)

	r, im := real(c), imag(c)
	fmt.Println(r, im)
}