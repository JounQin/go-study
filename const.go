package main

import (
	"fmt"
)

const (
	BIG   = 1 << 100
	SMALL = BIG >> 99
)

func needInt(i int) int {
	return i*10 + 1
}

func needFloat(f float64) float64 {
	return f * 0.1
}

func main() {
	fmt.Println(needInt(SMALL))
	fmt.Println(needFloat(SMALL))
	// fmt.Println(needInt(BIG))
	fmt.Println(needFloat(BIG))
}
