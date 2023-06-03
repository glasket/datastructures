package datastructures

import (
	"fmt"

	"github.com/glasket/datastructures/collection"
)

var _ collection.IImmutableIndexedCollection[int] = (*Tuple[int])(nil)

type Tuple[V comparable] struct {
	elements []V
}

func NewTuple[V comparable](values ...V) *Tuple[V] {
	return &Tuple[V]{
		elements: values,
	}
}

func (t *Tuple[V]) Contains(v V) bool {
	for _, e := range t.elements {
		if e == v {
			return true
		}
	}
	return false
}

func (t *Tuple[V]) Count() int {
	return len(t.elements)
}

func (t *Tuple[V]) IsEmpty() bool {
	return t.Count() == 0
}

func (t *Tuple[V]) String() string {
	return fmt.Sprintf("Tuple[%v]", t.elements)
}

func (t *Tuple[V]) IndexOf(v V) (int, error) {
	for i, e := range t.elements {
		if e == v {
			return i, nil
		}
	}
	return -1, fmt.Errorf("value %v not found", v)
}

func (t *Tuple[V]) Get(i int) (V, error) {
	if err := t.checkBounds(i); err != nil {
		return *new(V), err
	}
	return t.elements[i], nil
}

func (t *Tuple[V]) checkBounds(i int) error {
	if i < 0 || i >= t.Count() {
		return fmt.Errorf("index %d out of bounds", i)
	}
	return nil
}
