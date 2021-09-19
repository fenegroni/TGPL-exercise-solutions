package exercise7_7

import (
	"TGPL-exercise-solutions/chapter7/exercise7.8/music"
	"fmt"
	"sort"
	"testing"
	"time"
)

var tracks = []*music.Track{
	{"Go", "Delilah", "From the Roots Up", 2012, length("3m38s")},
	{"Go", "Moby", "Moby", 1992, length("3m37s")},
	{"Go Ahead", "Alicia Keys", "As I Am", 2007, length("4m36s")},
	{"Ready 2 Go", "Martin Solveig", "Smash", 2011, length("4m24s")},
}

func length(s string) time.Duration {
	d, err := time.ParseDuration(s)
	if err != nil {
		panic(s)
	}
	return d
}

func TestPrintTracks(t *testing.T) {
	fmt.Println("Click Artist once:")
	sort.Sort(music.ByArtist(tracks))
	music.PrintTracks(tracks)
	fmt.Println("Click Artist twice:")
	sort.Sort(sort.Reverse(music.ByArtist(tracks)))
	music.PrintTracks(tracks)
	fmt.Println("Click Year once:")
	sort.Sort(music.ByYear(tracks))
	music.PrintTracks(tracks)
	fmt.Println("Click Year twice:")
	sort.Sort(sort.Reverse(music.ByYear(tracks)))
	music.PrintTracks(tracks)
	fmt.Println("Custom sorting: Title, Year, Length:")
	sort.Sort(music.CustomSort{T: tracks, CombiLess: func(x, y *music.Track) bool {
		if x.Title != y.Title {
			return x.Title < y.Title
		}
		if x.Year != y.Year {
			return x.Year < y.Year
		}
		if x.Length != y.Length {
			return x.Length < y.Length
		}
		return false
	}})
	music.PrintTracks(tracks)
}
