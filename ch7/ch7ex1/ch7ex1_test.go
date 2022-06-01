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
	if _, err := c.Write([]byte("hello")); err != nil {
		t.Fatal("Unexpected error")
	}
	want = 5
	if got := int(c); got != want {
		t.Errorf("c is %d, want %d", got, want)
	}
	if _, err := c.Write([]byte("Dolly")); err != nil {
		t.Fatal("Unexpected error")
	}
	want = 10
	if got := int(c); got != want {
		t.Errorf("c is %d, want %d", got, want)
	}
	c = 0 // reset the counter
	var name = "Dolly"
	if _, err := fmt.Fprintf(&c, "hello, %s", name); err != nil {
		t.Fatal("Unexpected error")
	}
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
	_, err := w.Write([]byte("hello, dolly"))
	if err != nil {
		t.Fatal("Unexpected error")
	}
	want += 2
	if got := int(w); got != want {
		t.Errorf("w is %d, want %d", got, want)
	}
	_, err = w.Write([]byte("using a - counts as a word"))
	if err != nil {
		t.Fatal("Unexpected error")
	}
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
	_, err := l.Write([]byte("hello\ndolly"))
	if err != nil {
		t.Fatal("Unexpected error")
	}
	want += 2
	if got := int(l); got != want {
		t.Errorf("l is %d, want %d", got, want)
	}
	_, err = l.Write([]byte("using \n at the end does not count as a new line\n"))
	if err != nil {
		t.Fatal("Unexpected error")
	}
	want += 2
	if got := int(l); got != want {
		t.Errorf("l is %d, want %d", got, want)
	}
}
