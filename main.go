package main

import (
	"goMagic/core"
	"goMagic/downloader"
	"log"

	"github.com/PuerkitoBio/goquery"
)

type Article struct {
	Title string
	URL   string
}

type ToutiaoProcessor struct {
}

func (tt *ToutiaoProcessor) Process(p *downloader.Page) {
	q, err := p.Parser()

	if err != nil {
		log.Println(err)
		return
	}
	q.Find(".post").Each(func(index int, s *goquery.Selection) {
		content := s.Find(".content .title a")
		title := content.Text()
		url, _ := content.Attr("href")
		p.PutField("Title", title)
		p.PutField("URL", url)
		if url != "" {
			p.AddTargetURL(url)
		}
	})
	// var articles []Article
	// p.Objects(&articles)
}

func main() {
	core.NewMagic("test", &ToutiaoProcessor{}).AddURL("http://toutiao.io/"). /*.SetOutMode(pipe.MAPS)*/ Run()
}
