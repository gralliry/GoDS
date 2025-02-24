package atomicset

import (
	"sync"
	"sync/atomic"
)

type AtomicSet[K comparable] struct {
	maps sync.Map
	size atomic.Int64
}

func (s *AtomicSet[K]) Exist(key K) bool {
	_, ok := s.maps.Load(key)
	return ok
}

func (s *AtomicSet[K]) Insert(key K) bool {
	_, loaded := s.maps.LoadOrStore(key, struct{}{})
	if !loaded {
		s.size.Add(1)
	}
	return loaded
}

func (s *AtomicSet[K]) Remove(key K) bool {
	_, loaded := s.maps.LoadAndDelete(key)
	if loaded {
		s.size.Add(-1)
	}
	return loaded
}

func (s *AtomicSet[K]) Clear() {
	s.maps.Clear()
	s.size.Store(0)
}

func (s *AtomicSet[K]) Len() int {
	return int(s.size.Load())
}

func (s *AtomicSet[K]) Range(f func(key K) bool) {
	s.maps.Range(func(key, value any) bool {
		return f(key.(K))
	})
}
