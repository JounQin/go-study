package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	arr := make([][]uint8, dy)
	for i := range arr {
		arr[i] = make([]uint8, dx)
		for j := range arr[i] {
			arr[i][j] = uint8(i + j)
		}
	}
	return arr
}

func main() {
	pic.Show(Pic)
}
