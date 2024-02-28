package sub

import "github.com/maxguuse/100GoMistakePractice/7-ReturningInterfaces/client"

type InMemoryStore struct {
	storage []string
}

func NewInMemoryStore() client.Storer {
	return &InMemoryStore{}
}

func (s *InMemoryStore) Store(value string) {
	s.storage = append(s.storage, value)
}
