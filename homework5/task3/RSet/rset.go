package rset

import (
	"sync"
)

type Set struct {
	sync.RWMutex
	mm map[int]struct{}
}

func (s *Set) Add(i int) {
	s.Lock()
	s.mm[i] = struct{}{}
	s.Unlock()
}

func (s *Set) Has(i int) bool {
	s.RLock()
	defer s.RUnlock()
	_, ok := s.mm[i]
	return ok
}

func NewRSet() *Set {
	return &Set{
		mm: map[int]struct{}{},
	}
}
