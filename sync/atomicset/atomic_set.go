package atomicset

import (
	"sync"
)

type AtomicSet[K comparable] struct {
	maps sync.Map
}

func (s *AtomicSet[K]) Exist(key K) (ok bool) {
	_, ok = s.maps.Load(key)
	return ok
}

func (s *AtomicSet[K]) Insert(key K) (exist bool) {
	_, loaded := s.maps.LoadOrStore(key, struct{}{})
	return loaded
}

func (s *AtomicSet[K]) Remove(key K) (exist bool) {
	_, loaded := s.maps.LoadAndDelete(key)
	return loaded
}

func (s *AtomicSet[K]) Clear() {
	s.maps.Clear()
}

func (s *AtomicSet[K]) Range(f func(key K) bool) {
	s.maps.Range(func(key, value any) bool {
		return f(key.(K))
	})
}
