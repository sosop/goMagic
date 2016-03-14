package pipe

import (
	"io"
	"os"
)

type Pipeline interface {
	Out(s string) error
}

type ConsolePipeline struct {
	w io.Writer
}

func NewConsolePipeline() *ConsolePipeline {
	return &ConsolePipeline{os.Stdout}
}

func (console *ConsolePipeline) Out(s string) error {
	_, err := console.w.Write([]byte(s))
	return err
}
