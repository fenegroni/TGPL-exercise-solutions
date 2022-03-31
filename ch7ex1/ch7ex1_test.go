package ch7ex1

import (
	"fmt"
	"testing"
)

func TestByteCounter(t *testing.T) {
	var c ByteCounter
	want := 0
	if got := int(c); got != want {
		t.Fatalf("ByteCounter default value is %d, want %d", got, want)
	}
	c.Write([]byte("hello"))
	want = 5
	if got := int(c); got != want {
		t.Errorf("c is %d, want %d", got, want)
	}
	c.Write([]byte("Dolly"))
	want = 10
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
		t.Fatalf("WordCounter default value is %d, want %d", got, want)
	}
	w.Write([]byte("hello, dolly"))
	want += 2
	if got := int(w); got != want {
		t.Errorf("w is %d, want %d", got, want)
	}
	w.Write([]byte("using a - counts as a word"))
	want += 7
	if got := int(w); got != want {
		t.Errorf("w is %d, want %d", got, want)
	}
}

func TestLineCounter(t *testing.T) {
	var l LineCounter
	want := 0
	if got := int(l); got != want {
		t.Fatalf("LineCounter default value is %d, want %d", got, want)
	}
	l.Write([]byte("hello\ndolly"))
	want += 2
	if got := int(l); got != want {
		t.Errorf("l is %d, want %d", got, want)
	}
	l.Write([]byte("using \n at the end does not count as a new line\n"))
	want += 2
	if got := int(l); got != want {
		t.Errorf("l is %d, want %d", got, want)
	}
}
