package ch5ex18

import (
	"log"
	"os"
	"strings"
	"testing"
)

func Test_fetch(t *testing.T) {
	output := new(strings.Builder)
	log.SetOutput(output)
	url := "https://www.google.com"
	filename, _, err := Fetch(url)
	if err != nil && filename == "" {
		t.Fatalf("Could not fetch %q", url)
	}
	if !strings.Contains(output.String(), "deferred close call") {
		t.Errorf("Did not detect call to Close()")
	}
	_ = os.Remove(filename)
}
