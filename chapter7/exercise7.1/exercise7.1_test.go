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
