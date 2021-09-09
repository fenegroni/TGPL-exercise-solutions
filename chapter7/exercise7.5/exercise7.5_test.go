package exercise7_5

import (
	"io"
	"testing"
)

type zero struct{}

func (*zero) Read(p []byte) (n int, err error) {
	for i := 0; i < len(p); i++ {
		p[i] = 0
	}
	return len(p), nil
}

func TestLimitReader(t *testing.T) {
	tests := []struct {
		limit int64
	}{
		{0},
		{1},
		{2},
		{1000},
	}
	var z zero
	for _, test := range tests {
		eofZero := io.LimitReader(&z, test.limit)
		p := make([]byte, test.limit+10)
		n, err := eofZero.Read(p)
		n2, err2 := eofZero.Read(p)
		if int64(n) > test.limit || (err != io.EOF && err != nil) {
			t.Errorf("first pass io.LimitReader(%d).Read(%d) returns (%d, %v) want (%d, %v)", test.limit, len(p), n, err, test.limit, io.EOF)
		}
		if int64(n2) > 0 || err2 != io.EOF {
			t.Errorf("second pass io.LimitReader(%d).Read(%d) returns (%d, %v) want (%d, %v)", test.limit, len(p), n, err, 0, io.EOF)
		}
	}
}
