package main

import (
	"fmt"
	"sync"
)

type Set struct {
	mx   *sync.RWMutex
	data map[string]struct{}
}

func NewSet() *Set {
	return &Set{
		mx:   &sync.RWMutex{},
		data: make(map[string]struct{}),
	}
}

func (s *Set) Put(value string) error {
	s.mx.Lock()
	defer s.mx.Unlock()

	if s.IsExist(value) {
		return fmt.Errorf("value %s already exists", value)
	}
	s.data[value] = struct{}{}

	return nil
}

func (s *Set) GetAllData() ([]string, error) {
	s.mx.RLock()
	defer s.mx.RUnlock()

	list := make([]string, 0, len(s.data))
	for v := range s.data {
		list = append(list, v)
	}

	return list, nil
}

func (s *Set) IsExist(value string) bool {
	s.mx.RLock()
	defer s.mx.RUnlock()

	_, exist := s.data[value]
	return exist
}

func (s *Set) Remove(value string) {
	s.mx.Lock()
	defer s.mx.Unlock()

	delete(s.data, value)
}
