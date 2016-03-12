package downloader

import (
	"os"
	"testing"
)

var page *Page

func TestMain(m *testing.M) {
	page = NewPage("http://www.playhot.club")
	status := m.Run()
	os.Exit(status)
}

func TestParser(t *testing.T) {
	if _, err := page.Parser(); err != nil {
		t.Fatal(err)
	}
}
