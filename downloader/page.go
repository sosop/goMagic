package downloader

import (
	"encoding/json"
	"goMagic/scheduler"

	"github.com/PuerkitoBio/goquery"
)

type Page struct {
	URL    string
	Fields map[string][]string
	queue  scheduler.Queue
}

func NewPage(url string, queue scheduler.Queue) *Page {
	return &Page{url, make(map[string][]string, 32), queue}
}

func (p *Page) AddTargetURL(url string) {
	p.queue.Push(url)
}

func (p *Page) Parser() (*goquery.Document, error) {
	return goquery.NewDocument(p.URL)
}

func (p *Page) PutField(key, value string) {
	p.Fields[key] = append(p.Fields[key], value)
}

func (p *Page) Maps() []map[string]string {
	cols := len(p.Fields)
	colNames := make([]string, cols)
	count := 0
	for k, _ := range p.Fields {
		colNames[count] = k
		count++
	}

	if count == 0 {
		return nil
	}

	rows := len(p.Fields[colNames[0]])
	objs := make([]map[string]string, rows)
	for ir := 0; ir < rows; ir++ {
		obj := make(map[string]string, cols)
		for _, name := range colNames {
			obj[name] = p.Fields[name][ir]
		}
		objs[ir] = obj
	}

	return objs
}

func (p *Page) Objects(target interface{}) error {
	data, err := json.Marshal(p.Maps())
	if err != nil {
		return err
	}
	err = json.Unmarshal(data, target)
	return err
}
