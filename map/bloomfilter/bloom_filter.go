package bloomfilter

import (
	"GoDS/map/bitmap"
)

type BloomFilter[K any] struct {
	hashFunc func(v K) int
	maps     *bitmap.Bitmap
	cap      int
}

func New[K any](hashFunc func(v K) int, cap int) *BloomFilter[K] {
	bf := &BloomFilter[K]{
		hashFunc: hashFunc,
		maps:     bitmap.New(cap),
	}
	bf.cap = bf.maps.Cap()
	return bf
}

func (b *BloomFilter[K]) Add(v K) {
	for i := 0; i < b.cap; i++ {
		b.maps.Set(b.hashFunc(v) % b.cap)
	}
}

func (b *BloomFilter[K]) Check(v K) bool {
	return b.maps.Exist(b.hashFunc(v) % b.cap)
}

func (b *BloomFilter[K]) Clear() {
	b.maps.Clear()
}
