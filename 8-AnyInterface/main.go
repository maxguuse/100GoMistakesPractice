package main

import "log"

type Store struct {
	store []any
}

func (s *Store) Get(id int) any {
	return s.store[id]
}

func (s *Store) Store(value any) {
	s.store = append(s.store, value)
}

type Storage struct {
	stringStore []string
	intStore    []int
}

func (s *Storage) GetString(id int) string {
	return s.stringStore[id]
}

func (s *Storage) StoreString(value string) {
	s.stringStore = append(s.stringStore, value)
}

func (s *Storage) GetInt(id int) int {
	return s.intStore[id]
}

func (s *Storage) StoreInt(value int) {
	s.intStore = append(s.intStore, value)
}

type StringStorage interface {
	GetString(int) string
	StoreString(string)
}

type IntStorage interface {
	GetInt(int) int
	StoreInt(int)
}

func main() {
	s := &Store{}
	s.Store("hello")
	r := s.Get(0) // How the fuck would I know that I get a string here?

	s2 := &Storage{} // Simple as that
	s2.StoreString("hello")
	s2.StoreInt(42)

	// If you need some abstraction, use interfaces
	var ss StringStorage = s2
	r2 := ss.GetString(0)

	var is IntStorage = s2
	r3 := is.GetInt(0)

	log.Println(r)
	log.Println(r2)
	log.Println(r3)
}
