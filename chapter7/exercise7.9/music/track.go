package music

import (
	"fmt"
	"os"
	"text/tabwriter"
	"time"
)

type Track struct {
	Title  string
	Artist string
	Album  string
	Year   int
	Length time.Duration
}

type ByArtist []*Track

func (x ByArtist) Len() int           { return len(x) }
func (x ByArtist) Less(i, j int) bool { return x[i].Artist < x[j].Artist }
func (x ByArtist) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

type ByYear []*Track

func (x ByYear) Len() int           { return len(x) }
func (x ByYear) Less(i, j int) bool { return x[i].Year < x[j].Year }
func (x ByYear) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

type CustomSort struct {
	t    []*Track
	less func(x, y *Track) bool
}

func (x CustomSort) Len() int           { return len(x.t) }
func (x CustomSort) Less(i, j int) bool { return x.less(x.t[i], x.t[j]) }
func (x CustomSort) Swap(i, j int)      { x.t[i], x.t[j] = x.t[j], x.t[i] }

func PrintTracks(tracks []*Track) {
	const format = "%v\t%v\t%v\t%v\t%v\t\n"
	tw := new(tabwriter.Writer).Init(os.Stdout, 0, 8, 2, ' ', 0)
	_, _ = fmt.Fprintf(tw, format, "Title", "Artist", "Album", "Year", "Length")
	_, _ = fmt.Fprintf(tw, format, "-----", "------", "-----", "----", "------")
	for _, t := range tracks {
		_, _ = fmt.Fprintf(tw, format, t.Title, t.Artist, t.Album, t.Year, t.Length)
	}
	_ = tw.Flush()
}

func PrintTracksHTML(tracks []*Track) {
	//html/template
}
