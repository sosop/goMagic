package processor

import (
	"goMagic/downloader"
)

type Processor interface {
	Process(p *downloader.Page)
}
