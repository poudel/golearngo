package main

import (
	"fmt"
	"runtime"
)

/*
A `switch` statement is a shorter way to write a sequence of `if - else`
statements. It runs the first case whose value is equal to the condition
expression.

Switch cases evaluate cases from top to bottom, stopping when a case
succeeds.
*/

func main () {
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd
		// plan9, windows...
		fmt.Printf("%s.", os)
	}
}
