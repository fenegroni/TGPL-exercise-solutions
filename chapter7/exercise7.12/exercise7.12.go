package exercise7_12

import (
	"fmt"
	"html/template"
	"net/http"
	"strings"
)

type dollars float32
type database map[string]dollars

func (d dollars) String() string {
	return fmt.Sprintf("$%.2f", d)
}

func (db database) listHandler(resp http.ResponseWriter, _ *http.Request) {
	const tpl = `
<table>
  <tr>
    <th>Item</th>
    <th>Price</th>
  </tr>
  {{ range $index, $_ := . }}
  <tr>
    <td id="item{{ $index }}">{{ . }}</td>
    <td id="price{{ $index }}">{{ . }}</td>
  </tr>
  {{ end }}
</table>`
	t, err := template.New("db list").Parse(tpl)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(resp, "%s", err)
		return
	}
	var output strings.Builder
	err = t.Execute(&output, db)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(resp, "%s", err)
		return
	}
	fmt.Fprintln(resp, output.String())
}
