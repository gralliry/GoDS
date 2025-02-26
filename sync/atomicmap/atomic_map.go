// Package atomicmap This class simply puts a type restriction on sync.Map, without any explicit type conversion
package atomicmap

import "sync"

type AtomicMap[K any, V any] struct {
	maps sync.Map
}

func (m *AtomicMap[K, V]) Load(key K) (value V, ok bool) {
	return m.maps.Load(key)
}

func (m *AtomicMap[K, V]) LoadAndDelete(key K) (value V, loaded bool) {
	return m.maps.LoadAndDelete(key)
}

func (m *AtomicMap[K, V]) LoadOrStore(key K, value V) (actual V, loaded bool) {
	return m.maps.LoadOrStore(key, value)
}

func (m *AtomicMap[K, V]) CompareAndSwap(key K, old V, new V) (swapped bool) {
	return m.maps.CompareAndSwap(key, old, new)
}

func (m *AtomicMap[K, V]) CompareAndDelete(key K, old V) (deleted bool) {
	return m.maps.CompareAndDelete(key, old)
}

func (m *AtomicMap[K, V]) Store(key K, value V) {
	m.maps.Store(key, value)
}

func (m *AtomicMap[K, V]) Delete(key K) {
	m.maps.Delete(key)
}

func (m *AtomicMap[K, V]) Range(f func(key K, value V) bool) {
	m.maps.Range(func(key any, value any) bool {
		return f(key, value)
	})
}

func (m *AtomicMap[K, V]) Clear() {
	m.maps.Clear()
}
