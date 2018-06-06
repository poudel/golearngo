package main

import "fmt"

func main () {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	// among the init; condition; post
	// the init and post statements are optional
	bum := 1
	for bum < 1000 {
		bum += bum
	}
	fmt.Println(bum)

}
