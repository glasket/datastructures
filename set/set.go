/*
 * Copyright 2023, Christian Sigmon <cws@glasket.com>
 *
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/.
 */
package set

import (
	"encoding/json"
	"fmt"
)

var fil struct{} = struct{}{}

type Set[V comparable] struct {
	values map[V]struct{}
}

func NewSet[V comparable](size int) Set[V] {
	return Set[V]{
		values: make(map[V]struct{}, size),
	}
}

// Creates a set from a preexisting slice
//
// Ignores the Insert error for repeated values.
// If you want to catch errors, either use NewSet and manually Insert,
// or check the size of the slice against the resulting set.
func NewSetFromSlice[V comparable](slice []V) Set[V] {
	set := Set[V]{
		values: make(map[V]struct{}, len(slice)),
	}
	for _, v := range slice {
		set.Insert(v) // We don't care if it errors
	}
	return set
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
	for v := range s.values {
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

func (s Set[V]) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.Values())
}

// Unmarshal a JSON array into a Set.
//
// Uses NewSetFromSlice, but will return an error if
// the slice length doesn't match the set count.
func (s *Set[V]) UnmarshalJSON(data []byte) error {
	var slice []V
	err := json.Unmarshal(data, &slice)
	if err != nil {
		return err
	}
	*s = NewSetFromSlice(slice)
	if s.Count() != len(slice) {
		return fmt.Errorf("duplicates were present")
	}
	return nil
}
