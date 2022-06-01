package plot

import (
	"bytes"
	"encoding/xml"
	"testing"
)

// TestSurface validates the output of Surface is a valid SVG ready
// to be embedded in an HTTP response with content type SVG.
func TestSurface(t *testing.T) {
	buf := new(bytes.Buffer)
	written, err := Surface(buf, func(x, y float64) float64 {
		return 0
	})
	if err != nil {
		t.Fatalf("Surface errored: %s", err)
	}
	t.Logf("Surface wrote %d bytes.", written)
	t.Logf("Surface XML: %s", buf.String())
	tok, err := xml.NewDecoder(buf).Token()
	if err != nil {
		t.Fatalf("Surface errored: %s", err)
	}
	if start, ok := tok.(xml.StartElement); ok {
		if start.Name.Local != "svg" || start.Name.Space != "http://www.w3.org/2000/svg" {
			t.Fatal("Surface output is not SVG")
		}
	}
}
