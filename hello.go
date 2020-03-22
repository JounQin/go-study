package main

import (
	"fmt"
)

func add(x, y int) int {
	return x + y
}

func swap(x, y string) (string, string) {
	return y, x
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func main() {
	a, b := swap("a", "b")
	fmt.Println("Hello World!", add(1, 2), a, b)
	fmt.Println(swap("hello", "world"))
	fmt.Println(split(17))
}
