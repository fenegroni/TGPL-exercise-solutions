package ch7ex2

import "io"

type Writer struct {
	w io.Writer
	c int64
}

func CountingWriter(w io.Writer) (io.Writer, *int64) {
	ww := new(Writer)
	ww.w = w
	return ww, &ww.c
}

func (w *Writer) Write(p []byte) (n int, err error) {
	w.c += int64(len(p))
	return w.w.Write(p)
}
