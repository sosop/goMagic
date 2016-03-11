package processor

import (
	"goMagic/downloader"
)

var Processors map[string]Processor = make(map[string]Processor, 16)

type Processor interface {
	Process(p *downloader.Page)
}

func Add(processorName string, processor Processor) {
	Processors[processorName] = processor
}
