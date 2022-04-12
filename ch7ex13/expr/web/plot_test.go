package web

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

// TestPlot validates the output of Plot is a valid SVG response.
// Validating the SVG itself is done in TestSurface
func TestPlot(t *testing.T) {
	resp := httptest.NewRecorder()
	req := httptest.NewRequest("", "/?expr="+url.QueryEscape("1+2"), nil)
	Plot(resp, req)
	wantCode := http.StatusOK
	if got := resp.Result().StatusCode; got != wantCode {
		t.Fatalf("Response code: got %d, want %d", got, wantCode)
	}
	wantType := "image/svg+xml"
	if got := resp.Result().Header.Get("Content-type"); got != wantType {
		t.Errorf("Response Content-Type: got %q, want %q", got, wantType)
	}
}

// ExamplePlot runs an http server so you can experiment with different expressions.
func ExamplePlot() {
	// TODO implement ExamplePlot
}
