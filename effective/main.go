package main

import (
	"fmt"
	"log"
)

type ByteSize float64

const (
	_           = iota // ignore first value by assigning to blank identifier
	KB ByteSize = 1 << (10 * iota)
	MB
	GB
	TB
	PB
	EB
	ZB
	YB
)

func (b ByteSize) String() string {
	switch {
	case b >= YB:
		return fmt.Sprintf("%.2fYB", b/YB)
	case b >= ZB:
		return fmt.Sprintf("%.2fZB", b/ZB)
	case b >= EB:
		return fmt.Sprintf("%.2fEB", b/EB)
	case b >= PB:
		return fmt.Sprintf("%.2fPB", b/PB)
	case b >= TB:
		return fmt.Sprintf("%.2fTB", b/TB)
	case b >= GB:
		return fmt.Sprintf("%.2fGB", b/GB)
	case b >= MB:
		return fmt.Sprintf("%.2fMB", b/MB)
	case b >= KB:
		return fmt.Sprintf("%.2fKB", b/KB)
	}
	return fmt.Sprintf("%.2fB", b)
}

type ByteSlice []byte

func (p *ByteSlice) Write(data []byte) (n int, err error) {
	slice := *p
	slice = Append(slice, data...)
	*p = slice
	return len(data), nil
}

func Append[T any](slice []T, elements ...T) []T {
	l := len(slice)
	if l+len(elements) > cap(slice) {
		newSlice := make([]T, len(slice), l + len(elements))
		copy(newSlice, slice)
		slice = newSlice
	}
	slice = slice[:l+len(elements)]
	copy(slice[l:], elements)
	return slice
}

func Map(tz string) int {
	timeZone := map[string]int{
		"UTC": 0 * 60 * 60,
		"EST": -5 * 60 * 60,
		"CST": -6 * 60 * 60,
		"MST": -7 * 60 * 60,
		"PST": -8 * 60 * 60,
	}

	seconds, ok := timeZone[tz]

	if ok {
		return seconds
	}

	log.Println("unknown time zone:", tz)

	return 0
}

func main() {
	const (
		Enone  = 2
		Eio    = 6
		Einval = 10
	)
	a := [...]string{Enone: "no error", Eio: "Eio", Einval: "invalid argument"}
	s := []string{Enone: "no error", Eio: "Eio", Einval: "invalid argument"}
	m := map[int]string{Enone: "no error", Eio: "Eio", Einval: "invalid argument"}
	fmt.Println(a, s, m)

	var p = new([]int)
	fmt.Println(p)
	*p = append(*p, 10, 100, 1000)
	fmt.Println(p)

	var v = make([]int, 10, 100)
	fmt.Println(v)
	fmt.Println(len(v))
	fmt.Println(cap(v))

	fmt.Println(1<<63 - 1)
	fmt.Println(int(^uint(0) >> 1))
}
