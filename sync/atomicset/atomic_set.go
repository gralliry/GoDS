package atomicset

import (
	"sync"
)

type Set[K comparable] struct {
	maps sync.Map
}

func (s *Set[K]) Exist(key K) (ok bool) {
	_, ok = s.maps.Load(key)
	return ok
}

func (s *Set[K]) Insert(key K) (exist bool) {
	_, loaded := s.maps.LoadOrStore(key, struct{}{})
	return loaded
}

func (s *Set[K]) Remove(key K) (exist bool) {
	_, loaded := s.maps.LoadAndDelete(key)
	return loaded
}

func (s *Set[K]) Clear() {
	s.maps.Clear()
}

func (s *Set[K]) Range(f func(key K) bool) {
	s.maps.Range(func(key, value any) bool {
		return f(key.(K))
	})
}
