package goodsub

type InMemoryStore struct {
	storage []string
}

func NewInMemoryStore() *InMemoryStore {
	return &InMemoryStore{}
}

func (s *InMemoryStore) Store(value string) {
	s.storage = append(s.storage, value)
}
