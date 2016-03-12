package downloader

import "testing"

func BenchmarkParser(b *testing.B) {
	page := NewPage("http://www.playhot.club")
	if _, err := page.Parser(); err != nil {
		b.Fatal(err)
	}
}
