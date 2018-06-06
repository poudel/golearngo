package main

import (
	"fmt"
	"github.com/poudel/estring"
)


var c, python, java bool


func add (x, y int) int {
	return x + y
}


func swap (x, y string) (string, string) {
	// return values can be multiple
	return y, x
}


func split(sum int) (x, y int) {
	// naked return
	// named return values are returned automagically
	// naked return statements come with a readability cost
	x = sum * 4 / 9
	y = sum - x
	return
}


func main () {
	fmt.Println(add(10, 11))
	fmt.Println(swap("Hello", "World"))
	fmt.Println(split(10))
	estring.Estring()

	var i int
	fmt.Println(i, c, python, java)

	// var with initializers
	var j, k = 1, 2
	fmt.Println("Var with initializers:", j, k)

	// using the := construct
	m, n, o := true, false, "no!"
	fmt.Println("Using := construct", m, n, o)


	// basic go types
	/*
bool

string

int int8 int16 int32 int64
uint uint8 uint16 uint32 uint64 uintptr

byte // alias for uint8

rune // alias for int32
	// represents a Unicode code point

float32 float64

complex64 complex128

*/

	// type conversion
	// more simpler syntax
	// ii := 42
	// f := float64(ii)
	// u := uint(f)

	// type inference
	v := 42
	fmt.Printf("v is of type %T\n", v)

}
