package scheduler

const (
	Default_Max = 128
)

type Queue interface {
	Length() int
	Push(v interface{})
	Pop() interface{}
}

type BlockQueue interface {
	Queue
}

type MemQueue struct {
	max       int
	container chan interface{}
}

func NewMemQueueWithMax(max int) *MemQueue {
	mq := &MemQueue{max: max}
	mq.container = make(chan interface{}, max)
	return mq
}

func NewMemQueue() *MemQueue {
	mq := &MemQueue{max: Default_Max}
	mq.container = make(chan interface{}, Default_Max)
	return mq
}

func (m *MemQueue) Length() int {
	return len(m.container)
}

func (m *MemQueue) Push(v interface{}) {
	m.container <- v
}

func (m *MemQueue) Pop() interface{} {
	return <-m.container
}
