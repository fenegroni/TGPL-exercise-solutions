package exercise7_1

import (
	"fmt"
	"testing"
)

func TestByteCounter(t *testing.T) {
	var c ByteCounter
	want := 0
	if got := int(c); got != want {
		t.Errorf("ByteCounter default value is %d, want %d", got, want)
	}
	c.Write([]byte("hello"))
	want = 5
	if got := int(c); got != want {
		t.Errorf("c is %d, want %d", got, want)
	}
	c = 0 // reset the counter
	var name = "Dolly"
	fmt.Fprintf(&c, "hello, %s", name)
	want = 12
	if got := int(c); got != want {
		t.Errorf("c is %d, want %d", got, want)
	}

}

func TestWordCounter(t *testing.T) {
	var w WordCounter
	want := 0
	if got := int(w); got != want {
		t.Errorf("WordCounter default value is %d, want %d", got, want)
	}
}

func TestLineCounter(t *testing.T) {
	var w LineCounter
	want := 0
	if got := int(w); got != want {
		t.Errorf("LineCounter default value is %d, want %d", got, want)
	}
}
