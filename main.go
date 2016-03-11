package main

import (
	"fmt"
	"goMagic/downloader"

	"github.com/PuerkitoBio/goquery"
)

type Article struct {
	Title string
	URL   string
}

func main() {
	p := downloader.NewPage("http://toutiao.io/")
	q, err := p.Parser()
	if err != nil {
		panic(err)
	}
	q.Find(".post").Each(func(index int, s *goquery.Selection) {
		content := s.Find(".content .title a")
		title := content.Text()
		url, _ := content.Attr("href")
		p.PutField("Title", title)
		p.PutField("URL", url)
	})
	var articles []Article
	p.Objects(&articles)
	fmt.Println(articles)
}