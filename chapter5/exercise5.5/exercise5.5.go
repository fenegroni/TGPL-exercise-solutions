// CountWordsAndImages counts the words and images in each HTML document.
// The URL to each document is supplied on the command line.
package exercise5_5

import (
	"bufio"
	"fmt"
	"golang.org/x/net/html"
	"net/http"
	"os"
	"strings"
)

func main() {
	for _, url := range os.Args[1:] {
		words, images, err := CountWordsAndImages(url)
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "exercise5.5: %v\n", err)
			continue
		}
		fmt.Printf("%s, words: %d, images: %d\n", url, words, images)
	}
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
	if n == nil {
		return
	}
	if !(n.Type == html.ElementNode && (n.Data == "script" || n.Data == "style")) {
		if n.Type == html.TextNode {
			words += countWords(n.Data)
		}
		if n.Type == html.ElementNode && n.Data == "img" {
			images++
		}
		w, i := countWordsAndImages(n.FirstChild)
		words += w
		images += i
	}
	w, i := countWordsAndImages(n.NextSibling)
	words += w
	images += i
	return
}

func countWords(input string) (words int) {
	scanner := bufio.NewScanner(strings.NewReader(input))
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		words++
	}
	return
}
