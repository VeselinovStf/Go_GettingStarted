package main

import "fmt"

const packageConst = 5.43 // Package level const

// Constant block
const (
	first = 1
	second = 2
)

// iota
const(
	third = iota
	fourth = iota
)

const(
	a = iota + 1 // Constant expression
	b = 2 << iota
	c = 2 << iota
	d = 2 << iota
	e
	f
)

func main() {
	// Declare and Initialize at a same time
	// value must be determ at compile-time

	const i = 3 // Implicid type const

	fmt.Println(i)

	fmt.Println(i + 1)
	fmt.Println(i + 1.4)

	const x int = 3

	fmt.Println(x + 1)
	fmt.Println(float32(x) + 1.4)

	//iota
	fmt.Println(third,fourth)

	fmt.Println(a,b,c,d,e,f)
}