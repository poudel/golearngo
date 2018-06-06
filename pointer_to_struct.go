package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{1, 2}
	p := &v

	fmt.Println(
		"We could also write (*p).X but the language allows",
		"to write p.X without the explicit dereference",
	)
	p.X = 100
	fmt.Println(v)
}
