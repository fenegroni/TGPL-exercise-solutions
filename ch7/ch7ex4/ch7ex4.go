package ch7ex4

import (
	"io"
)

type Reader struct {
	s string
}

func (r *Reader) WriteTo(io.Writer) (n int64, err error) {
	panic("implement me")
}

func (r *Reader) Seek(int64, int) (int64, error) {
	panic("implement me")
}

func (r *Reader) UnreadByte() error {
	panic("implement me")
}

func (r *Reader) ReadByte() (byte, error) {
	panic("implement me")
}

func (r *Reader) ReadAt([]byte, int64) (n int, err error) {
	panic("implement me")
}

func (r *Reader) Read(p []byte) (n int, err error) {
	if len(p) == 0 {
		return 0, nil
	}
	if len(r.s) == 0 {
		return 0, io.EOF
	}
	for n = 0; n < len(p); n++ {
		if n >= len(r.s) {
			return n, io.EOF
		}
		p[n] = r.s[n]
	}
	return n, nil
}

func NewReader(s string) *Reader {
	r := new(Reader)
	r.s = s
	return r
}
