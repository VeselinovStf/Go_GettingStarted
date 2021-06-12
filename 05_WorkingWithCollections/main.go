package main

import (
	"fmt"
)

func main() {
	// Arrays - fix type of collection with simular type elements
	var arr [3]int
	arr[0] = 1
	arr[1] = 2
	arr[2] = 3
	fmt.Println(arr)
	fmt.Println(arr[1])

	arrImplicid := [4]int{99, 66, 77, 88}
	fmt.Println(arrImplicid)

	// Slices - build on top of Arrays
	sliceFromArray := arrImplicid[:] // From start : to end
	fmt.Println(sliceFromArray)

	// Slice points to array values
	arrImplicid[0] = 9999
	fmt.Println(sliceFromArray)

	sliceFromArray[1] = 1111
	fmt.Println(sliceFromArray)
	fmt.Println(arrImplicid)

	// Init slice
	slice := []int{0, 1, 2, 3}
	fmt.Println(slice)

	// Add to slice - Dynamic resize
	slice = append(slice, 4)
	fmt.Println(slice)

	s1 := slice[1:]
	s2 := slice[:2]
	s3 := slice[1:2]
	fmt.Println(s1, s2, s3)

	// Maps - Key-Value
	m := map[string]int{"A": 1, "B": 2}
	fmt.Println(m)
	fmt.Println(m["A"])

	delete(m, "A")
	fmt.Println(m["A"])
	fmt.Println(m["B"])

	//struct- collection of any type of data
	type user struct {
		Id        int
		FirstName string
		LastName  string
	}
	var u user

	fmt.Println(u)
	u.Id = 1
	u.FirstName = "Gosho"
	u.LastName = "Peshov"
	fmt.Println(u)

	u2 := user{Id: 2, FirstName: "Misho", LastName: "Mishov"}
	fmt.Println(u2)

}
