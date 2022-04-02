package ch5ex18

import (
	"io"
	"log"
	"net/http"
	"os"
	"path"
)

func Fetch(url string) (string, int64, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", 0, err
	}
	defer resp.Body.Close()
	local := path.Base(resp.Request.URL.Path)
	if local == "/" || local == "." {
		local = "index.html"
	}
	f, err := os.Create(local)
	if err != nil {
		return "", 0, err
	}
	n, err := io.Copy(f, resp.Body)
	defer func() {
		// Close file, but prefer error from Copy, if any.
		if closeErr := f.Close(); err == nil {
			err = closeErr
		}
		log.Println("deferred close call")
	}()
	return local, n, err
}
