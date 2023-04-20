package main

import (
	"testing"
	"unicode/utf8"
)

func TestReverse(t *testing.T) {
	testcases := []struct {
		in, want string
	}{
		{"Hello, world", "dlrow ,olleH"},
		{" ", " "},
		{"!12345", "54321!"},
	}
	for _, tc := range testcases {
		rev, err := Reverse(tc.in)
		if err != nil {
			t.Errorf("Reverse: %q, err: %v", tc.in, err)
		}
		if rev != tc.want {
			t.Errorf("Reverse: %q, want %q", rev, tc.want)
		}
	}
}

func FuzzReverse(f *testing.F) {
	testcases := []string{"Hello, world", " ", "!12345"}
	for _, tc := range testcases {
		f.Add(tc) // Use f.Add to provide a seed corpus
	}
	f.Fuzz(func(t *testing.T, orig string) {
		rev, err := Reverse(orig)
		if err != nil {
			return
		}
		doubleRev, err := Reverse(rev)
		if err != nil {
			return
		}
		t.Logf("Number of runes: orig=%d, rev=%d, doubleRev=%d",
			utf8.RuneCountInString(orig),
			utf8.RuneCountInString(rev),
			utf8.RuneCountInString(doubleRev),
		)
		if orig != doubleRev {
			t.Errorf("Before: %q, after: %q", orig, doubleRev)
		}
		if utf8.ValidString(orig) && !utf8.ValidString(rev) {
			t.Errorf("Reverse produced invalid UTF-8 string %q", rev)
		}
	})
}
