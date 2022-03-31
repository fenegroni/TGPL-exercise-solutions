package ch7ex4

import (
	"io"
)

type Reader struct {
	s string
}

func (r *Reader) WriteTo(w io.Writer) (n int64, err error) {
	panic("implement me")
}

func (r *Reader) Seek(offset int64, whence int) (int64, error) {
	panic("implement me")
}

func (r *Reader) UnreadByte() error {
	panic("implement me")
}

func (r *Reader) ReadByte() (byte, error) {
	panic("implement me")
}

func (r *Reader) ReadAt(p []byte, off int64) (n int, err error) {
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
