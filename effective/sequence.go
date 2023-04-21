package main

import (
	"fmt"
	"sort"
)

type Integer interface {
	uint |
		uintptr |
		uint8 |
		uint16 |
		uint32 |
		uint64 |
		int |
		int8 |
		int16 |
		int32 |
		int64
}

type Number interface {
	Integer |
		float32 |
		float64
}

type Sequence[T Number] []T

func (s Sequence[T]) Len() int {
	return len(s)
}

func (s Sequence[T]) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s Sequence[T]) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s Sequence[T]) Copy() Sequence[T] {
	return append(Sequence[T]{}, s...)
}

func (s Sequence[T]) String() string {
	s = s.Copy()
	sort.Sort(s)
	return fmt.Sprintf("%v", []T(s))
}

type IntSequence []int

func (s IntSequence) Copy() IntSequence {
	return append(IntSequence{}, s...)
}

func (s IntSequence) String() string {
	s = s.Copy()
	sort.IntSlice(s).Sort()
	return fmt.Sprintf("%v", []int(s))
}
