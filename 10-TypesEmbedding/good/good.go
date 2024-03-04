package good

import "sync"

type inMem struct {
	mu sync.Mutex
	m  map[string]int
}

func NewInMem() *inMem {
	return &inMem{
		m: map[string]int{},
	}
}

func (i *inMem) Add(key string, value int) {
	i.mu.Lock()
	defer i.mu.Unlock()
	i.m[key] = value
}

func (i *inMem) Get(key string) int {
	i.mu.Lock()
	defer i.mu.Unlock()
	return i.m[key]
}
