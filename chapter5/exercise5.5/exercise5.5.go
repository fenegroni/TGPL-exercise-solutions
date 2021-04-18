// CountWordsAndImages counts the words and images in each HTML document.
// The URL to each document is supplied on the command line.
package main

import (
	"fmt"
	"golang.org/x/net/html"
	"net/http"
)

func main() {

}

// CountWordsAndImages does an HTTP GET request for the HTML document url and
// returns the number of words and images in it
func CountWordsAndImages(url string) (words, images int, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	parseTree, err := html.Parse(resp.Body)
	_ = resp.Body.Close()
	if err != nil {
		err = fmt.Errorf("parsing HTML: %s", err)
		return
	}
	words, images = countWordsAndImages(parseTree)
	return
}

func countWordsAndImages(n *html.Node) (words, images int) {
	return
}
