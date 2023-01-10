package set

import (
	"fmt"
)

var fil struct{} = struct{}{}

type Set[V comparable] struct {
	values map[V]struct{}
}

func NewSet[V comparable]() *Set[V] {
	return &Set[V]{
		values: map[V]struct{}{},
	}
}

// Inserts a given value into the set.
//
// Throws an error if the value is already present.
func (s *Set[V]) Insert(value V) error {
	if s.Contains(value) {
		return fmt.Errorf("value %v is already present", value)
	}
	s.values[value] = fil
	return nil
}

func (s *Set[V]) Remove(value V) error {
	if !s.Contains(value) {
		return fmt.Errorf("value %v is not present", value)
	}
	delete(s.values, value)
	return nil
}

func (s *Set[V]) Clear() {
	for v, _ := range s.values {
		delete(s.values, v)
	}
}

func (s Set[V]) Contains(value V) bool {
	_, ok := s.values[value]
	return ok
}

func (s Set[V]) Count() int {
	return len(s.values)
}

func (s Set[V]) Values() []V {
	values := make([]V, len(s.values))
	i := 0

	for v := range s.values {
		values[i] = v
		i++
	}

	return values
}
