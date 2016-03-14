package pipe

import (
	"goMagic/downloader"
	"io"
	"os"
)

const (
	NORMAL = 1
	MAPS   = 2
)

type Pipeline interface {
	Out(p *downloader.Page) error
	Close() error
	Mode(mode int)
}

type ConsolePipeline struct {
	w    io.WriteCloser
	mode int
}

func NewConsolePipeline() *ConsolePipeline {
	return &ConsolePipeline{os.Stdout, NORMAL}
}

func (console *ConsolePipeline) Out(p *downloader.Page) error {
	var err error
	_, err = console.w.Write([]byte("crawler: " + p.URL + "\n\n"))
	switch console.mode {
	case NORMAL:
		for k, v := range p.Fields {
			_, err = console.w.Write([]byte(k + ": \n"))
			for _, v := range v {
				_, err = console.w.Write([]byte("\t" + v + "\n"))
			}
		}
	case MAPS:
		for _, v := range p.Maps() {
			for kk, vv := range v {
				_, err = console.w.Write([]byte(kk + ": " + vv + "\t\t"))
			}
			_, err = console.w.Write([]byte("\n"))
		}
	default:
		console.w.Write([]byte("choose out mode!"))
	}
	return err
}

func (console *ConsolePipeline) Close() error {
	return console.w.Close()
}

func (console *ConsolePipeline) Mode(mode int) {
	console.mode = mode
}

type FilePipeline struct {
	w *os.File
}

func NewFilePipeline(filename string) *FilePipeline {
	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	return &FilePipeline{file}
}

func (filePipeline *FilePipeline) Out(p *downloader.Page) error {
	_, err := filePipeline.w.WriteString("")
	return err
}

func (filePipeline *FilePipeline) Close() error {
	return filePipeline.w.Close()
}

func (filePipeline *FilePipeline) Mode(mode int) {
}
