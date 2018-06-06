package main

import "fmt"

/*
- Constants are similar to variables
- They are declared + assigned using the `const` keyword
- They cannot be declared using := syntax
- They can be character, boolean, or numeric values
*/

const Pi = 3.14

func main () {
	const World = " "

	fmt.Println("Hello", World)
	fmt.Println("Happy", Pi, "Day")

	const Truth = true
	fmt.Println("Go rules?", Truth)
}
