package exercise7_2

import (
	"os"
	"testing"
)

func TestCountingWriter(t *testing.T) {
	w, c := CountingWriter(os.Stdout)
	t.Logf("%v, %v", w, c)
}
