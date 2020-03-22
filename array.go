package main

import (
	"fmt"
)

func main() {
	var arr [2]string
	arr[0] = "Hello"
	arr[1] = "World"
	fmt.Println(arr[0], arr[1], arr)

	primes := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes, len(primes))
	sub := primes[2:5]
	fmt.Println(sub)

	sub[0] = 1
	fmt.Println(sub, primes)
}
