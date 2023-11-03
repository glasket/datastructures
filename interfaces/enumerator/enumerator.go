/*
 * Copyright 2023 Christian Sigmon <cws@glasket.com>
 *
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/.
 */

package enumerator

// IEnumerable is the base interface for all collections. It enforces that
// a structure must provider an enumerator and a way to get a slice of values.
type IEnumerable[V any] interface {
	GetEnumerator() IEnumerator[V]
	Values() []V
}

// IEnumerator provides the functionality to enumerate over a collection.
//
// Implementations should be thread safe, with the enumerator acting on a
// slice copy of the collection.
type IEnumerator[V any] interface {
	Current() V
	Next() bool
	Reset()
}

type enumerator[V any] struct {
	elements []V
	position int
	current  V
}

type enumerable[V any] struct {
	elements []V
}

func (e *enumerable[V]) GetEnumerator() IEnumerator[V] {
	return &enumerator[V]{
		elements: e.elements,
		position: -1,
		current:  *new(V),
	}
}

func (e *enumerable[V]) Values() []V {
	return e.elements
}

// GetSliceEnumerable gets an enumerable for the built-in slice type.
func GetSliceEnumerable[V any](elements []V) IEnumerable[V] {
	return &enumerable[V]{
		elements: elements,
	}
}

// GetMapKeyEnumerable gets an enumerable for the built-in map type's keys.
func GetMapKeyEnumerable[K comparable, V any](elements map[K]V) IEnumerable[K] {
	keys := make([]K, 0)
	for key := range elements {
		keys = append(keys, key)
	}
	return &enumerable[K]{
		elements: keys,
	}
}

// GetMapValueEnumerable gets an enumerable for the built-in map type's values.
func GetMapValueEnumerable[K comparable, V any](elements map[K]V) IEnumerable[V] {
	values := make([]V, 0)
	for _, value := range elements {
		values = append(values, value)
	}
	return &enumerable[V]{
		elements: values,
	}
}

// GetMapValueEnumerable gets an enumerable for the built-in map type's keys and values.
// The pair is stored in a [2]any array.
func GetMapEnumerable[K comparable, V any](elements map[K]V) IEnumerable[[2]any] {
	pairs := make([][2]any, 0)
	for key, value := range elements {
		pairs = append(pairs, [2]any{key, value})
	}
	return &enumerable[[2]any]{
		elements: pairs,
	}
}

func (e *enumerator[V]) Current() V {
	return e.current
}

func (e *enumerator[V]) Next() bool {
	if e.position+1 >= len(e.elements) {
		e.current = *new(V)
		return false
	}

	e.position += 1
	e.current = e.elements[e.position]
	return true
}

func (e *enumerator[V]) Reset() {
	e.position = -1
	e.current = *new(V)
}
