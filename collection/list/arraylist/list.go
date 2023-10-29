/*
 * Copyright 2023 Christian Sigmon <cws@glasket.com>
 *
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/.
 */

package arraylist

import (
	"fmt"

	"github.com/glasket/datastructures/collection/list"
	"github.com/glasket/datastructures/interfaces/enumerator"
)

var _ list.IList[int] = (*List[int])(nil)

// TODO Documentation

type List[V comparable] struct {
	elements []V
}

func New[V comparable]() *List[V] {
	return &List[V]{
		elements: make([]V, 0),
	}
}

func NewFromSlice[V comparable](l []V) *List[V] {
	return &List[V]{
		elements: l,
	}
}

func (l *List[V]) Add(v V) {
	l.elements = append(l.elements, v)
}

func (l *List[V]) Remove(v V) {
	i, err := l.IndexOf(v)
	if err != nil {
		return
	}
	l.elements = append(l.elements[:i], l.elements[i+1:]...)
}

func (l *List[V]) InsertAt(i int, v V) error {
	if err := l.checkBounds(i); err != nil {
		if i == len(l.elements) {
			l.Add(v)
			return nil
		}
		return err
	}
	l.elements = append(l.elements[:i], append([]V{v}, l.elements[i:]...)...)
	return nil
}

func (l *List[V]) RemoveAt(i int) error {
	if err := l.checkBounds(i); err != nil {
		return err
	}
	l.elements = append(l.elements[:i], l.elements[i+1:]...)
	return nil
}

func (l *List[V]) Get(i int) (V, error) {
	if err := l.checkBounds(i); err != nil {
		return *new(V), err
	}
	return l.elements[i], nil
}

func (l *List[V]) Set(i int, v V) error {
	if err := l.checkBounds(i); err != nil {
		return err
	}
	l.elements[i] = v
	return nil
}

func (l *List[V]) Count() int {
	return len(l.elements)
}

func (l *List[V]) ToSlice() []V {
	return l.elements
}

func (l *List[V]) Clear() {
	l.elements = make([]V, 0)
}

func (l *List[V]) Contains(v V) bool {
	for _, e := range l.elements {
		if e == v {
			return true
		}
	}
	return false
}

func (l *List[V]) String() string {
	return fmt.Sprintf("List[%v]", l.elements)
}

func (l *List[V]) IsEmpty() bool {
	return l.Count() == 0
}

func (l *List[V]) IndexOf(v V) (int, error) {
	for i, e := range l.elements {
		if e == v {
			return i, nil
		}
	}
	return -1, fmt.Errorf("value %v not found", v)
}

func (l *List[V]) GetEnumerator() enumerator.IEnumerator[V] {
	return enumerator.GetSliceEnumerable(l.elements).GetEnumerator()
}

func (l *List[V]) checkBounds(i int) error {
	if i < 0 || i >= len(l.elements) {
		return fmt.Errorf("index %d out of bounds", i)
	}
	return nil
}
