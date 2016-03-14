package pipe

import "testing"

func TestConsolePipelineOut(t *testing.T) {
	pipeline := NewConsolePipeline()
	defer pipeline.Close()
	if err := pipeline.Out("testing console"); err != nil {
		t.Fatal(err)
	}
}

func TestFilePipelineOut(t *testing.T) {
	pipeline := NewFilePipeline("/Users/mac/filePipeline")
	defer pipeline.Close()
	if err := pipeline.Out("testing file"); err != nil {
		t.Fatal(err)
	}
}
