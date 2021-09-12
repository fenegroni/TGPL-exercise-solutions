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
		eofZero := LimitReader(&z, test.limit)
		p := make([]byte, test.limit/2+10)
		n, err := eofZero.Read(p)
		n2, err2 := eofZero.Read(p)
		n3, err3 := eofZero.Read(p)
		if int64(n+n2+n3) > test.limit {
			t.Errorf("LimitReader(%d).Read(%d) reads %d bytes, want %d", test.limit, len(p), n+n2+n3, test.limit)
		}
		if err != io.EOF && err != nil {
			t.Errorf("LimitReader(%d).Read: first pass reading %d bytes returns error %v want either nil or io.EOF", test.limit, len(p), err)
		}
		if err2 != io.EOF && err2 != nil {
			t.Errorf("LimitReader(%d).Read: second pass reading %d bytes returns error %v want either nil or io.EOF", test.limit, len(p), err)
		}
		if err3 != io.EOF {
			t.Errorf("LimitReader(%d).Read: third pass reading %d bytes returns error %v want io.EOF", test.limit, len(p), err)
		}
	}
}
