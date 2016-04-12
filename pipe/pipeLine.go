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
	return CommonOut(console.w, p, console.mode)
}

func (console *ConsolePipeline) Close() error {
	return console.w.Close()
}

func (console *ConsolePipeline) Mode(mode int) {
	console.mode = mode
}

type FilePipeline struct {
	w    *os.File
	mode int
}

func NewFilePipeline(filename string) *FilePipeline {
	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	return &FilePipeline{w: file}
}

func (filePipeline *FilePipeline) Out(p *downloader.Page) error {
	return CommonOut(filePipeline.w, p, filePipeline.mode)
}

func (filePipeline *FilePipeline) Close() error {
	return filePipeline.w.Close()
}

func (filePipeline *FilePipeline) Mode(mode int) {
	filePipeline.mode = mode
}

func CommonOut(w io.Writer, p *downloader.Page, mode int) error {
	var err error
	_, err = w.Write([]byte("crawler: " + p.URL + "\n\n"))
	switch mode {
	case NORMAL:
		for k, v := range p.Fields {
			_, err = w.Write([]byte(k + ": \n"))
			for _, v := range v {
				_, err = w.Write([]byte("\t" + v + "\n"))
			}
		}
	case MAPS:
		for _, v := range p.Maps() {
			for kk, vv := range v {
				_, err = w.Write([]byte(kk + ": " + vv + "\t\t"))
			}
			_, err = w.Write([]byte("\n"))
		}
	default:
		w.Write([]byte("choose out mode!"))
	}
	return err
}
