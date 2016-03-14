package pipe

import "testing"

var pipeline Pipeline

func TestConsolePipelineOut(t *testing.T) {
	pipeline = NewConsolePipeline()
	if err := pipeline.Out("testing console"); err != nil {
		t.Fatal(err)
	}
}
