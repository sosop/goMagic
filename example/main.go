package main

import (
	"goMagic/core"
	"goMagic/downloader"
	"goMagic/pipe"
	"log"
	"strings"

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
		url = strings.TrimSpace(url)
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
	core.NewMagic("test", &ToutiaoProcessor{}).AddURL("http://toutiao.io/"). /*.SetThread(8).SetPipeline(pipeline).SetQueue(q)*/ SetOutMode(pipe.MAPS).Run()
}
