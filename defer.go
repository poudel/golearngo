package main

import "fmt"

func main () {
	// defer waits execution until
	defer fmt.Println("World!")
	fmt.Println("Hello")
}
