package main

func main() {

	// Conditional Looping
	var i int
	for i < 5 {
		// Short-circuit
		if i == 2 {
			println("Its Two")
			// Break From Loop
			break
		}

		// Build in print line
		println(i)
		i++

		if i == 1 {
			// Continue to next loop iteration
			continue
		}
	}

	// Post clouse - Classic for
	j := 50 // or use implicit init in for loop
	for ; j <= 55; j++ {
		println(j)
	}

	//Ifinity
	for {
		if j == 60 {
			break
		}
		println(j)
		j++
	}

	// Loop over collections
	slice := []int{1, 2, 3}
	for _, element := range slice {
		println(element)
	}
}
