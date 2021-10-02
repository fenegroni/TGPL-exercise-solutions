package exercise7_9

import (
	"TGPL-exercise-solutions/chapter7/exercise7.9/music"
	"fmt"
	"os"
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
	htmlString, err := music.PrintTracksAsHTMLString(tracks)
	if err != nil {
		t.Fatalf("PrintTracksAsHTMLString error: %v", err)
	}
	f, _ := os.Create("index.html")
	fmt.Fprintln(f, htmlString)
	_ = f.Close()
}
