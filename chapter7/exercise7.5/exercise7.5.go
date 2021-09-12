package exercise7_5

import "io"

func LimitReader(r io.Reader, limit int64) io.Reader {
	return r
}
