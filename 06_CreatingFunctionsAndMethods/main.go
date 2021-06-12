package main

import (
	"errors"
	"fmt"
)

func main() {
	// Function Invocation
	printHello()
	message := "Hello"
	printString(message)
	printTwoString(message, "Second Message")
	
	returnStatement := sayTry()
	fmt.Println(returnStatement )
	
	errorStatement:= getError()
	fmt.Println(errorStatement)
	
	code, error := getCodedError()
	fmt.Println(code,error )
	
	// Write-only variable
	c, _ := getCodedError()
	fmt.Println(c )
}

// Function Declaration
func printHello(){
	fmt.Println("Hello, playground")
}

// Pass Function Parameter
func printString(message string){
	fmt.Println(message)
}

// Same argument types - simple type deffinition
func printTwoString(message, messageTwo string){
	fmt.Println(message,messageTwo)
}

// Return from funtion
func sayTry() bool{
	return true
}

// Return error from funtion
func getError() error{
	return errors.New("Something Went Whrong")
}

// Return more then one result
func getCodedError() (int, error){
	return 0, errors.New("Something Went Whrong")
}