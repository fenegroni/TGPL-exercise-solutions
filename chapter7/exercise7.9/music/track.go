package music

import (
	"fmt"
	"html/template"
	"os"
	"strings"
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

type ByTitle []*Track

func (x ByTitle) Len() int           { return len(x) }
func (x ByTitle) Less(i, j int) bool { return x[i].Title < x[j].Title }
func (x ByTitle) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

type ByArtist []*Track

func (x ByArtist) Len() int           { return len(x) }
func (x ByArtist) Less(i, j int) bool { return x[i].Artist < x[j].Artist }
func (x ByArtist) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

type ByAlbum []*Track

func (x ByAlbum) Len() int           { return len(x) }
func (x ByAlbum) Less(i, j int) bool { return x[i].Album < x[j].Album }
func (x ByAlbum) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

type ByYear []*Track

func (x ByYear) Len() int           { return len(x) }
func (x ByYear) Less(i, j int) bool { return x[i].Year < x[j].Year }
func (x ByYear) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

type ByLength []*Track

func (x ByLength) Len() int           { return len(x) }
func (x ByLength) Less(i, j int) bool { return x[i].Length < x[j].Length }
func (x ByLength) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

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

func PrintTracksAsHTMLString(tracks []*Track) (HTMLString string, err error) {
	const tpl = `
<table>
  <tr>
    <th><a id="HeaderLink0" href="?sor=Title">Title</a></th>
    <th><a id="HeaderLink1" href="?sort=Artist">Artist</a></th>
    <th><a id="HeaderLink2" href="?sort=Album">Album</a></th>
    <th><a id="HeaderLink3" href="?sort=Year">Year</a></th>
    <th><a id="HeaderLink4" href="?sort=Length">Length</a></th>
  </tr>
  {{ range $index, $_ := . }}
  <tr id="row{{ $index }}">
    <td id="row{{ $index }}col0">{{ .Title }}</td>
    <td id="row{{ $index }}col1">{{ .Artist }}</td>
    <td id="row{{ $index }}col2">{{ .Album }}</td>
    <td id="row{{ $index }}col3">{{ .Year }}</td>
    <td id="row{{ $index }}col4">{{ .Length }}</td>
  </tr>
  {{ end }}
</table>`
	// FIXME Add the row and rowcol ids to the html
	var t *template.Template
	if t, err = template.New("webpage").Parse(tpl); err != nil {
		return "", err
	}
	var output strings.Builder
	if err = t.Execute(&output, tracks); err != nil {
		return "", err
	}
	return output.String(), nil
}
