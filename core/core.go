package core

import (
	"goMagic/downloader"
	"goMagic/pipe"
	proc "goMagic/processor"
	"goMagic/scheduler"
	"runtime"
)

var (
	defaultThread = runtime.NumCPU()
)

type Magic struct {
	TaskName string
	Proc     proc.Processor
	threadN  int
	queue    scheduler.Queue
	pipeline pipe.Pipeline
}

func NewMagic(taskName string, processor proc.Processor) *Magic {
	return &Magic{taskName, processor, defaultThread, scheduler.NewMemQueue(), pipe.NewConsolePipeline()}
}

func (m *Magic) SetQueue(q scheduler.Queue) *Magic {
	m.queue = q
	return m
}

func (m *Magic) SetPipeline(p pipe.Pipeline) *Magic {
	m.pipeline = p
	return m
}

func (m *Magic) SetOutMode(mode int) *Magic {
	m.pipeline.Mode(mode)
	return m
}

func (m *Magic) AddURL(url string) *Magic {
	m.queue.Push(url)
	return m
}

func (m *Magic) SetThread(nThread int) *Magic {
	m.threadN = nThread
	return m
}

func (m *Magic) execute() {
	url := m.queue.Pop()
	p := downloader.NewPage(url.(string))
	m.Proc.Process(p)
	m.pipeline.Out(p)
}

func (m *Magic) Run() {
	defer m.pipeline.Close()
	for m.queue.Length() > 0 {
		m.execute()
	}
}
