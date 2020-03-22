package main

import (
	"fmt"
)

func Sqrt(x float64) float64 {
	z := 1.0
	n := 10
	for i := 0; i < n; i++ {
		z -= (z*z - x) / (2 * z)
	}
	return z
}

func main() {
	fmt.Println(Sqrt(2))
}
