package music

import (
	"fmt"
	"sort"
	"time"
)

func ExamplePrintTracks() {
	length := func(s string) time.Duration {
		d, err := time.ParseDuration(s)
		if err != nil {
			panic(s)
		}
		return d
	}
	var tracks = []*Track{
		{"Go", "Delilah", "From the Roots Up", 2012, length("3m38s")},
		{"Go", "Moby", "Moby", 1992, length("3m37s")},
		{"Go Ahead", "Alicia Keys", "As I Am", 2007, length("4m36s")},
		{"Ready 2 Go", "Martin Solveig", "Smash", 2011, length("4m24s")},
	}
	fmt.Println("Click Artist once:")
	sort.Sort(ByArtist(tracks))
	PrintTracks(tracks)
	fmt.Println("Click Artist twice:")
	sort.Sort(sort.Reverse(ByArtist(tracks)))
	PrintTracks(tracks)
	fmt.Println("Click Year once:")
	sort.Sort(ByYear(tracks))
	PrintTracks(tracks)
	fmt.Println("Click Year twice:")
	sort.Sort(sort.Reverse(ByYear(tracks)))
	PrintTracks(tracks)
	fmt.Println("Custom sorting: Title, Year, Length:")
	sort.Sort(CustomSort{T: tracks, CombiLess: func(x, y *Track) bool {
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
	PrintTracks(tracks)
}
