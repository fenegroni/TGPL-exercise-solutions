package ch7ex9

import (
	"errors"
	exercise58 "github.com/fenegroni/TGPL-exercise-solutions/ch5/ch5ex8"
	"github.com/fenegroni/TGPL-exercise-solutions/ch7ex8/stable"
	"github.com/fenegroni/TGPL-exercise-solutions/ch7ex9/music"
	"golang.org/x/net/html"
	"io"
	"net/http"
	"net/http/httptest"
	"sort"
	"strings"
	"testing"
	"time"
)

type trackTable []*music.Track

func (tracks trackTable) clickHandler(w http.ResponseWriter, r *http.Request) {
	var stableTracks stable.SortableIndexable
	switch r.URL.Query().Get("sort") {
	case "Title":
		stableTracks = music.ByTitle(tracks)
	case "Artist":
		stableTracks = music.ByArtist(tracks)
	case "Album":
		stableTracks = music.ByAlbum(tracks)
	case "Year":
		stableTracks = music.ByYear(tracks)
	case "Length":
		stableTracks = music.ByLength(tracks)
	}
	sort.Sort(stable.NewSorted(stableTracks))
	tracksAsHTML, _ := music.PrintTracksAsHTML(tracks)
	_, _ = w.Write([]byte(tracksAsHTML))
}

func TestPrintTracksAsHTML(t *testing.T) {
	length := func(s string) time.Duration {
		d, err := time.ParseDuration(s)
		if err != nil {
			panic(s)
		}
		return d
	}
	var tracks = trackTable{
		{"Go", "Delilah", "From the Roots Up", 2012, length("3m38s")},
		{"Go", "Moby", "Moby", 1992, length("3m37s")},
		{"Go Ahead", "Alicia Keys", "As I Am", 2007, length("4m36s")},
		{"Ready 2 Go", "Martin Solveig", "Smash", 2011, length("4m24s")},
	}
	var tracksOrderedByLengthTitle = trackTable{
		{"Go", "Moby", "Moby", 1992, length("3m37s")},
		{"Go", "Delilah", "From the Roots Up", 2012, length("3m38s")},
		{"Go Ahead", "Alicia Keys", "As I Am", 2007, length("4m36s")},
		{"Ready 2 Go", "Martin Solveig", "Smash", 2011, length("4m24s")},
	}
	wantHTML, err := music.PrintTracksAsHTML(tracksOrderedByLengthTitle)
	if err != nil {
		t.Fatalf("PrintTracksAsHTML error: %v", err)
	}
	tracksAsHTML, err := music.PrintTracksAsHTML(tracks)
	if err != nil {
		t.Fatal(err)
	}
	htmlSortedByLength, err := clickOnColumnHeader(tracksAsHTML, "Length", tracks.clickHandler)
	if err != nil {
		t.Fatal(err)
	}
	htmlSortedByLengthTitle, err := clickOnColumnHeader(htmlSortedByLength, "Title", tracks.clickHandler)
	if err != nil {
		t.Fatal(err)
	}
	if htmlSortedByLengthTitle != wantHTML {
		t.Fatal("htmlSortedByLengthTitle does not match wantHTML")
	}
}

func clickOnColumnHeader(htmlTable string, column string, handler http.HandlerFunc) (string, error) {
	docTracks, _ := html.Parse(strings.NewReader(htmlTable))
	linkSortBy, err := getHeaderLink("By"+column, docTracks)
	if err != nil {
		return "", err
	}
	sortRequest := httptest.NewRequest("", "/"+linkSortBy, nil)
	responseWriter := httptest.NewRecorder()
	handler(responseWriter, sortRequest)
	sortResponse := responseWriter.Result()
	body, err := io.ReadAll(sortResponse.Body)
	return string(body), err
}

func getHeaderLink(by string, doc *html.Node) (string, error) {
	node := exercise58.ElementByID(doc, "HeaderLink"+by)
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
