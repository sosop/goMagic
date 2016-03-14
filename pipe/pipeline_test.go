package pipe

import "testing"

var pipeline Pipeline

func TestConsolePipelineOut(t *testing.T) {
	pipeline = NewConsolePipeline()
	defer pipeline.Close()
	if err := pipeline.Out("testing console"); err != nil {
		t.Fatal(err)
	}
}
