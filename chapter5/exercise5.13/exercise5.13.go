package main

import (
	"fmt"
	"golang.org/x/net/html"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	// Crawl the web breadth-first,
	// starting from the command-line arguments.
	breadthFirst(crawl, os.Args[1:2])
}

func breadthFirst(f func(string) []string, worklist []string) {
	seen := make(map[string]bool)
	for len(worklist) > 0 {
		items := worklist
		worklist = nil
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				worklist = append(worklist, f(item)...)
			}
		}
	}
}

// crawl wraps calls to Extract, logging any errors,
// allowing breadthFirst to continue crawling through all the hyperlinks.
func crawl(url string) []string {
	list, err := Extract(url)
	if err != nil {
		log.Print(err)
	}
	return list
}

// Extract makes an HTTP GET request to the specified URL address, parses
// the response in HTML, and returns the links in the HTML document.
func Extract(address string) ([]string, error) {
	resp, err := http.Get(address)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("getting %q: %s", address, resp.Status)
	}
	doc, err := html.Parse(resp.Body)
	_ = resp.Body.Close()
	if err != nil {
		return nil, fmt.Errorf("parsing %q as HTML: %v", address, err)
	}
	var links []string
	// TODO only save links to urls within the selected domain
	// analyse path and determine if it ends with a name or a slash or anything
	// if the path is empty, it's index.html
	var folderpath, filepath string
	if resp.Request.URL.Path == "" {
		folderpath = resp.Request.URL.Hostname()+
		filepath = folderpath+"/index.html"
	} else if strings.LastIndex(resp.Request.URL.Path, "/") == len(resp.Request.URL.Path)-1 {
		folderpath = resp.Request.URL.Hostname()+"/"+resp.Request.URL.Path
		filepath = folderpath+"/index.html"
	} else {
		lastSlash := strings.LastIndex(resp.Request.URL.Path, "/")
		lastDot := strings.LastIndex(resp.Request.URL.Path, ".")
		if lastDot > lastSlash && lastDot < len(resp.Request.URL.Path)-1 {
			folderpath = resp.Request.URL.Hostname() + "/" + resp.Request.URL.Path[:lastSlash]
			filepath = folderpath+resp.Request.URL.Path[lastSlash:]
		}
	}
	// TODO check for '.', '..', '.exe', etc...
	// save file
	_ = os.MkdirAll(folderpath, 0)
	os.Create(filepath)
	visitNode := func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key != "href" {
					continue
				}
				link, err := resp.Request.URL.Parse(a.Val)
				if err != nil {
					continue
				}
				links = append(links, link.String())
			}
		}
	}
	forEachNode(doc, visitNode, nil)
	return links, nil
}

// forEachNode calls the functions pre(x) and post(x) for each node
// x in the tree rooted at n. Both functions are optional.
// pre is called before the children are visited (preorder) and
// post is called after (postorder).
func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}

	if post != nil {
		post(n)
	}
}
