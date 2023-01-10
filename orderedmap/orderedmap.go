/*
 * Copyright 2022, Christian Sigmon <cws@glasket.com>
 *
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/.
 */
package orderedmap

import "fmt"

// An OrderedMap maintains the insertion order of keys into a contained map.
type OrderedMap[K comparable, V any] struct {
	mapping map[K]V
	keys    []K
}

// Constructs a new OrderedMap with instantiated fields.
func NewOrderedMap[K comparable, V any](size int) OrderedMap[K, V] {
	return OrderedMap[K, V]{
		mapping: map[K]V{},
		keys:    make([]K, 0, size),
	}
}

// Returns the key slice of the OrderedMap
func (m OrderedMap[K, V]) Keys() []K {
	return m.keys
}

// Returns the assigned value at the given key, or an error if the key is not present in the map.
func (m OrderedMap[K, V]) Get(key K) (value V, err error) {
	var ok bool
	value, ok = m.mapping[key]
	if !ok {
		return value, fmt.Errorf("key %v was not present", key)
	}
	return value, nil
}

// Assigns the given value to the given key.
//
// The key maintains its prior position. To update the position then use SetAndUpdate.
func (m *OrderedMap[K, V]) Set(key K, value V) {
	if !m.Contains(key) {
		m.keys = append(m.keys, key)
	}
	m.mapping[key] = value
}

// Assigns the given value to the given key and shifts the key to the end of the order
func (m *OrderedMap[K, V]) SetAndUpdate(key K, value V) {
	if m.Contains(key) {
		for i, v := range m.keys {
			if key == v {
				l := len(m.keys)
				newKeys := make([]K, l)
				for j, k := 0, 0; j < l-1; j, k = j+1, k+1 {
					if j == i {
						k += 1
					}
					newKeys[j] = m.keys[k]
				}
				newKeys[l-1] = key
				m.keys = newKeys
				break
			}
		}
	} else {
		m.keys = append(m.keys, key)
	}
	m.mapping[key] = value
}

// Removes the given key and its assigned value.
//
// Returns an error if the key was not assigned.
func (m *OrderedMap[K, V]) Remove(key K) error {
	if !m.Contains(key) {
		return fmt.Errorf("key %v was not present", key)
	}
	delete(m.mapping, key)
	for i, v := range m.keys {
		if key == v {
			m.keys = append(m.keys[:i], m.keys[i+1:]...)
			break
		}
	}
	return nil
}

// Returns true if the given key is present in the underlying map, otherwise false.
func (m OrderedMap[K, V]) Contains(key K) bool {
	_, ok := m.mapping[key]
	return ok
}

// Returns the number of keys in the map.
func (m OrderedMap[K, V]) Count() int {
	return len(m.keys)
}
