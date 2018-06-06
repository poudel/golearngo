package main

import "fmt"

/*
- They are high-precision values
- An untyped constant takes the type needed by its context
*/

const (
	// create huge number by shifting a 1 bit of left 100 places.
	// in other words, the binary number that is 1 followed by 100 zeroes.
	Big = 1 << 100

	// shift it right again 99 places, so we end up with 1<<1 or 2.
	Small = Big >> 99
)


func needInt(x int) int {
	return x * 10 + 1
}

func needFloat (x float64) float64 {
	return x * 0.1
}


func main () {
	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
}