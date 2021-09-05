package exercise7_1

import (
	"fmt"
	"testing"
)

func TestByteCounter_defaultValue(t *testing.T) {
	var c ByteCounter
	want := "0"
	if got := fmt.Sprint(c); got != want {
		t.Errorf("c is %s, want %s", got, want)
	}
}

func TestByteCounter_Write(t *testing.T) {
	var c ByteCounter
	c.Write([]byte("hello"))
	want := "5"
	if got := fmt.Sprint(c); got != want {
		t.Errorf("c is %s, want %s", got, want)
	}
	c = 0 // reset the counter
	var name = "Dolly"
	fmt.Fprintf(&c, "hello, %s", name)
	want = "12"
	if got := fmt.Sprint(c); got != want {
		t.Errorf("c is %s, want %s", got, want)
	}
}
