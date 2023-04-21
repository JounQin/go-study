package main

import (
	"fmt"
	"testing"
)

func TestSequence_Copy(t *testing.T) {
	s := Sequence[int]{1, 2, 3}
	s2 := s.Copy()
	s2[0] = 4
	if s[0] == s2[0] {
		t.Errorf("Sequence.Copy() = %d; want %d", s2[0], s[0])
	}
	if len(s) != len(s2) {
		t.Errorf("len Sequence.Copy() = %d; want %d", len(s2), len(s))
	}
	if cap(s) != cap(s2) {
		t.Errorf("cap Sequence.Copy() = %d; want %d", cap(s2), cap(s))
	}
}

func TestSequence_String(t *testing.T) {
	fmt.Println(Sequence[int]{3, 1, 2})
}
