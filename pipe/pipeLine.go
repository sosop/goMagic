package pipe

import (
	"io"
	"os"
)

type Pipeline interface {
	Out(s string) error
	Close() error
}

type ConsolePipeline struct {
	w io.WriteCloser
}

func NewConsolePipeline() *ConsolePipeline {
	return &ConsolePipeline{os.Stdout}
}

func (console *ConsolePipeline) Out(s string) error {
	_, err := console.w.Write([]byte(s))
	return err
}

func (console *ConsolePipeline) Close() error {
	return console.w.Close()
}
