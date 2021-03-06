package ch7ex12

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
  {{ range $item, $price := . }}
  <tr>
    <td id="item">{{ $item }}</td>
    <td id="price">{{ $price }}</td>
  </tr>
  {{ end }}
</table>`
	t, err := template.New("db list").Parse(tpl)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		_, _ = fmt.Fprintf(resp, "%s", err)
		return
	}
	var output strings.Builder
	err = t.Execute(&output, db)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		_, _ = fmt.Fprintf(resp, "%s", err)
		return
	}
	_, _ = fmt.Fprintln(resp, output.String())
}
