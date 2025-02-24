package set

type Set[K comparable] struct {
	maps map[K]struct{}
}

func (s *Set[K]) Exist(key K) bool {
	_, ok := s.maps[key]
	return ok
}

func (s *Set[K]) Insert(key K) {
	s.maps[key] = struct{}{}
}

func (s *Set[K]) Remove(key K) {
	delete(s.maps, key)
}

func (s *Set[K]) Clear() {
	s.maps = make(map[K]struct{})
}

func (s *Set[K]) Len() int {
	return len(s.maps)
}

func (s *Set[K]) Range(f func(key K) bool) {
	for key := range s.maps {
		if !f(key) {
			break
		}
	}
}
