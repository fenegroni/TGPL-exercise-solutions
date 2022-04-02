package ch5ex13

import (
	"fmt"
	"golang.org/x/net/html"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
)

var DownloadDir string

// BreadthFirst calls f for each item in the worklist.
// Any items returned by f are added to the worklist.
// f is called at most once for each item.
func BreadthFirst(f func(string) []string, worklist []string) {
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

// Crawl wraps calls to Extract, logging any errors,
// allowing BreadthFirst to continue crawling through all the hyperlinks.
func Crawl(address string) []string {
	links, err := Extract(address)
	if err != nil {
		log.Printf("error: Extract(%q): %v", address, err)
		return nil
	}
	var sameDomainLinks []string
	addressUrl, _ := url.Parse(address)
	for _, a := range links {
		aUrl, _ := url.Parse(a)
		if aUrl.Host == addressUrl.Host {
			sameDomainLinks = append(sameDomainLinks, a)
		}
	}
	return sameDomainLinks
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
	var folderpath, filepath string
	requestUrl := resp.Request.URL
	folderpath = DownloadDir + "/" + requestUrl.Hostname() + "__" + requestUrl.Port()
	if requestUrl.Path == "" {
		filepath = folderpath + "/index.html"
	} else if strings.LastIndex(requestUrl.Path, "/") == len(requestUrl.Path)-1 {
		folderpath += "/" + requestUrl.Path
		filepath = folderpath + "/index.html"
	} else {
		lastSlash := strings.LastIndex(requestUrl.Path, "/")
		lastDot := strings.LastIndex(requestUrl.Path, ".")
		if lastDot > lastSlash && lastDot < len(requestUrl.Path)-1 {
			folderpath += "/" + requestUrl.Path[:lastSlash]
			filepath = folderpath + requestUrl.Path[lastSlash:]
		}
		// TODO validate what happens if the condition is not met
	}
	// TODO check for '.', '..', '.exe', etc...
	// TODO save file content
	err = os.MkdirAll(folderpath, 0777)
	if err != nil {
		return nil, fmt.Errorf("unable to create directory %q", folderpath)
	}
	saveFile, err := os.Create(filepath)
	if err != nil {
		return nil, fmt.Errorf("unable to create file %q", filepath)
	}
	_ = saveFile.Close()
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
