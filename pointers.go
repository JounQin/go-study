package main

import (
	"fmt"
)

func main() {
	i, j := 42, 2701
	p := &i

	fmt.Println(p)
	fmt.Println(*p)
	*p = *p / 2
	fmt.Println(i)

	p = &j

	fmt.Println(p)
	fmt.Println(j)
}
