package main

import "fmt"

func main () {
	primes := [6]int{3, 5, 7, 11, 13}

	var s []int = primes[1:4]
	fmt.Println(s)

	// slices are like references to arrays
	// changing the elements of a slice modifies the corresponding
	// elements of the underlying array
	// other slices that share the same underlying array will see the changes
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b)

	b[0] = "XXX"
	fmt.Println(a, b)
	fmt.Println(names)
}
