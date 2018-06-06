package main

import (
	"fmt"
	"math"
)

func sqrt (x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func pow (x, n, lim float64) float64 {
	// variables declared by the short statement are only in scope
	// until the end of the `if`
	// they are also available in relevant `else` blocks
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	return lim
}

func main () {
	fmt.Println(sqrt(2))
	fmt.Println(sqrt(-4))
	fmt.Println(
		"POWER",
		pow(3, 2, 10),
		pow(3, 3, 20),
	)
}
