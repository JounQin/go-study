package main

import (
	"fmt"
	"testing"
)

func TestByteSize(t *testing.T) {
	fmt.Println(YB)
	fmt.Println(ByteSize(1e13))
}

func TestByteSlice(t *testing.T) {
	var b ByteSlice
	_, _ = fmt.Fprintf(&b, "This hour has %d days\n", 7)
	if string(b) != "This hour has 7 days\n" {
		t.Errorf("ByteSlice() = %s; want `This hour has 7 days\n`", b)
	}
}

func FuzzByteSlice(f *testing.F) {
	f.Add("Hello World")
	f.Fuzz(func(t *testing.T, text string) {
		var b ByteSlice
		_, _ = fmt.Fprintf(&b, "%s", text)
		if string(b) != text {
			t.Errorf("ByteSlice() = %s; want %s", b, text)
		}
	})
}

func TestAppend(t *testing.T) {
	slice := []int{1, 2, 3}
	elements := []int{4, 5, 6}

	slice = Append(slice, elements...)

	if len(slice) != 6 {
		t.Errorf("Append() = %d; want 6", len(slice))
	}
}

func TestMap(t *testing.T) {
	seconds := Map("UTC")

	if seconds != 0 {
		t.Errorf("Map(\"UTC\") = %d; want 0", seconds)
	}
}
