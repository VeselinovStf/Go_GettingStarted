package main

import "fmt"

func main(){
	// init pointer - pointer opperator
	var firstName *string = new(string)

	// dereference - dereference opperator
	*firstName = "Pesho"

	// Print pointer address
	fmt.Println(firstName)

	// dereference - get data try pointer ( to what data this pointer points )
	fmt.Println(*firstName)

	lastName := "Goshov"
	fmt.Println(lastName)

	// address of opperator
	ptr := &lastName
	fmt.Println(ptr)
	fmt.Println(ptr, *ptr)

	// change value of variable - address is same
	lastName = "Ivanov"
	fmt.Println(ptr, *ptr)
}