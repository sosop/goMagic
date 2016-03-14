package core

import (
	"goMagic/downloader"
	"goMagic/pipe"
	"goMagic/processor"
	"goMagic/scheduler"
)

type Magic struct {
}

func NewMagic() *Magic {
	return &Magic{}
}

func (m *Magic) Add(name, url string, proc *processor.Processor, pipeline *pipe.Pipeline, queue *scheduler.Queue) *Magic {
	downloader.URLs[name] = url
	downloader.Processors[name] = proc
	downloader.Pipelines[name] = pipeline
	downloader.Queues[name] = queue
	return m
}

func (m *Magic) AddDefault(name, url string, proc *processor.Processor) *Magic {
	m.Add(name, url, proc, pipe.NewConsolePipeline(), scheduler.NewMemQueue())
	return m
}

func (m *Magic) StartThread(nThread int) *Magic {
	return m
}

func (m *Magic) execute() {
	for name, url := range downloader.URLs {
		downloader.Queues[name].Push(url)
	}
}

func (m *Magic) Run() {

}
