package scheduler

import (
	"os"
	"testing"
)

var bq BlockQueue

func TestMain(m *testing.M) {
	bq = NewMemQueue()
	status := m.Run()
	os.Exit(status)
}

func TestLength(t *testing.T) {
	bq.Push(1)
	if bq.Length() != 1 {
		t.Fatal("push error")
	}
	bq.Pop()
	if bq.Length() != 0 {
		t.Fatal("pop error")
	}
}
