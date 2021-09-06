package exercise7_2

import (
	"fmt"
	"strings"
	"testing"
)

func TestCountingWriter(t *testing.T) {
	str1, str2, str3 := "First string", " second string", " third string"
	sb := new(strings.Builder)
	cw, c := CountingWriter(sb)
	want, _ := fmt.Fprint(cw, str1)
	if *c != int64(want) {
		t.Errorf("Write %q: want %d, got %d", str1, want, *c)
	}
	wantmore, _ := fmt.Fprint(cw, str2)
	want += wantmore
	if *c != int64(want) {
		t.Errorf("Write %q to the same CountingWriter, want %d, got %d", str2, want, *c)
	}
	cw2, c2 := CountingWriter(sb)
	want2, _ := fmt.Fprint(cw2, str3)
	if *c2 != int64(want2) && *c != int64(want) {
		t.Errorf("Write %q to a different CountingWriter, using the same strings.Builder: want %d, got %d", str3, want2, *c2)
	}
	strWant := str1 + str2 + str3
	if got := sb.String(); got != strWant {
		t.Errorf("String built so far: want %q, got %q", strWant, got)
	}
}
