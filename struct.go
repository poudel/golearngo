package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main () {
	v := Vertex{1, 2}
	fmt.Println(v)

	fmt.Println("Struct fields are accessed using dot (.)")

	fmt.Println("Vertex.X ->", v.X, "...", "Vertex.Y ->", v.Y)
}
