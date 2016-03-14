package downloader

import (
	"encoding/json"
	"goMagic/pipe"
	proc "goMagic/processor"
	sche "goMagic/scheduler"

	"github.com/PuerkitoBio/goquery"
)

var Pipelines = make(map[string]*pipe.Pipeline, 16)
var Processors = make(map[string]*proc.Processor, 16)
var Queues = make(map[string]*sche.Queue, 16)
var URLs = make(map[string]string, 16)

type Page struct {
	URL    string
	Fields map[string][]string
}

func NewPage(url string) *Page {
	return &Page{url, make(map[string][]string, 32)}
}

func (p *Page) Parser() (*goquery.Document, error) {
	return goquery.NewDocument(p.URL)
}

func (p *Page) PutField(key, value string) {
	p.Fields[key] = append(p.Fields[key], value)
}

func (p *Page) Objects(target interface{}) error {
	cols := len(p.Fields)
	colNames := make([]string, cols)
	count := 0
	for k, _ := range p.Fields {
		colNames[count] = k
		count++
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

	data, err := json.Marshal(objs)
	if err != nil {
		return err
	}
	err = json.Unmarshal(data, target)
	return err
}
