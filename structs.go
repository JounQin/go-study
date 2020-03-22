package main

import (
	"fmt"
)

type Vertex struct {
	x, y int
}

func main() {
	v := Vertex{2, 3}
	fmt.Println(v)
	v.x = 4
	fmt.Println(v)
	p := &v
	p.x = 1e9
	fmt.Println(v)
}
