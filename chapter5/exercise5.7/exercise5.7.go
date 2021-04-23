// Findlinks prints the links in an HTML document read from standard input.
package main

import (
	"fmt"
	"golang.org/x/net/html"
	"log"
	"net/http"
	"os"
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

	resp, err := http.Get("http://www.google.com")
	if err != nil {
		log.Fatalf("Can't get url %s: %v", "http://www.google.com", err)
	}
	parseTree, err := html.Parse(resp.Body)
	_ = resp.Body.Close()
	if err != nil {
		log.Fatalf("parsing HTML: %v", err)
	}
	for _, link := range Visit(nil, parseTree) {
		fmt.Println(link)
	}
}

func ForEachNode(n *html.Node, pre, post func(n *html.Node)) {
	//TODO implement ForEachNode pg.133
}

// Visit appends to links each link found in n and returns the result.
func Visit(links []string, n *html.Node) []string {
	if n == nil {
		return links
	}
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}
	return Visit(Visit(links, n.FirstChild), n.NextSibling)
}
