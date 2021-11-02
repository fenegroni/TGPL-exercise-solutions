package exercise7_9

import (
	exercise5_8 "TGPL-exercise-solutions/chapter5/exercise5.8"
	"TGPL-exercise-solutions/chapter7/exercise7.9/music"
	"errors"
	"fmt"
	"golang.org/x/net/html"
	"net/http/httptest"
	"os"
	"sort"
	"strings"
	"testing"
	"time"
)

func TestPrintTracksHTML(t *testing.T) {
	length := func(s string) time.Duration {
		d, err := time.ParseDuration(s)
		if err != nil {
			panic(s)
		}
		return d
	}
	var tracks = []*music.Track{
		{"Go", "Delilah", "From the Roots Up", 2012, length("3m38s")},
		{"Go", "Moby", "Moby", 1992, length("3m37s")},
		{"Go Ahead", "Alicia Keys", "As I Am", 2007, length("4m36s")},
		{"Ready 2 Go", "Martin Solveig", "Smash", 2011, length("4m24s")},
	}
	var tracksByLengthTitle = []*music.Track{
		{"Go", "Moby", "Moby", 1992, length("3m37s")},
		{"Go", "Delilah", "From the Roots Up", 2012, length("3m38s")},
		{"Go Ahead", "Alicia Keys", "As I Am", 2007, length("4m36s")},
		{"Ready 2 Go", "Martin Solveig", "Smash", 2011, length("4m24s")},
	}
	htmlTracksByLengthTitle, err := music.PrintTracksAsHTMLString(tracksByLengthTitle)
	if err != nil {
		t.Fatalf("PrintTracksAsHTMLString error: %v", err)
	}
	docTracksByLengtAndTitle, err := html.Parse(strings.NewReader(htmlTracksByLengthTitle))
	if err != nil {
		t.Fatal(err)
	}
}

func clickOnColumnHeader(column string, tracks []*music.Track) (string, error) {
	htmlTracks, err := music.PrintTracksAsHTMLString(tracks)
	if err != nil {
		return "", err
	}
	docTracks, _ := html.Parse(strings.NewReader(htmlTracks))
	linkSortBy, err := getHeaderLink("By"+column, docTracks)
	if err != nil {
		return "", err
	}
	// FIXME I want to now specify a type that can be used in a web server as an HTTP handler
	//  for requests within a webmusic module that provides this facility.
	sortBy := httptest.NewRequest("", "/"+linkSortBy, nil).URL.Query().Get("sort")
	if sortBy == "" {
		return "", fmt.Errorf("no sort key in header link %q", linkSortBy)
	}
	if sortBy != column {
		return "", fmt.Errorf("%q is not a valid sort value for column %q", sortBy, column)
	}
	// FIXME Use my implementation of stable sorting from ex7.8
	switch sortBy {
	case "Length":
		sort.Stable(music.ByLength(tracks))
	}
	return music.PrintTracksAsHTMLString(tracks)
}

func getHeaderLink(by string, doc *html.Node) (string, error) {
	node := exercise5_8.ElementByID(doc, "HeaderLink"+by)
	if node == nil {
		return "", errors.New("no element HeaderLink" + by)
	}
	linkText := ""
	for _, a := range node.Attr {
		if a.Key == "href" {
			linkText = a.Val
		}
	}
	return linkText, nil
}
