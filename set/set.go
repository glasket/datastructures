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

type Set[V comparable] map[V]struct{}

func NewSet[V comparable](size int) Set[V] {
	return make(map[V]struct{}, size)
}

// Creates a set from a preexisting slice
//
// Ignores the Insert error for repeated values.
// If you want to catch errors, either use NewSet and manually Insert,
// or check the size of the slice against the resulting set.
func NewSetFromSlice[V comparable](slice []V) Set[V] {
	set := NewSet[V](len(slice))
	for _, v := range slice {
		set.Insert(v) // We don't care if it errors
	}
	return set
}

// Inserts a given value into the set.
//
// Throws an error if the value is already present.
func (s Set[V]) Insert(value V) error {
	if s.Contains(value) {
		return fmt.Errorf("value %v is already present", value)
	}
	s[value] = fil
	return nil
}

// Removes the given value from the set.
//
// Throws an error if the value is not present.
func (s Set[V]) Remove(value V) error {
	if !s.Contains(value) {
		return fmt.Errorf("value %v is not present", value)
	}
	delete(s, value)
	return nil
}

// Deletes all values from the set.
func (s Set[V]) Clear() {
	for v := range s {
		delete(s, v)
	}
}

// Returns true if the given value is present in the set.
func (s Set[V]) Contains(value V) bool {
	_, ok := s[value]
	return ok
}

// Returns the number of values in the set.
func (s Set[V]) Count() int {
	return len(s)
}

// Returns a slice of all values in the set.
func (s Set[V]) Values() []V {
	values := make([]V, len(s))
	i := 0

	for v := range s {
		values[i] = v
		i++
	}

	return values
}

// Test if two sets are equal.
//
// Two sets are equal if they contain the same values.
func (s Set[V]) Equals(other Set[V]) bool {
	if s.Count() != other.Count() {
		return false
	}

	for v := range s {
		if !other.Contains(v) {
			return false
		}
	}
	return true
}

// Returns a new set containing the values in both sets.
//
// Does not modify either set.
func (s Set[V]) Union(other Set[V]) Set[V] {
	union := NewSet[V](s.Count() + other.Count())

	for v := range s {
		union.Insert(v)
	}
	for v := range other {
		union.Insert(v)
	}

	return union
}

// Returns a new set containing the values in both sets.
//
// Does not modify either set.
func (s Set[V]) Intersection(other Set[V]) Set[V] {
	intersection := NewSet[V](s.Count())

	for v := range s {
		if other.Contains(v) {
			intersection.Insert(v)
		}
	}

	return intersection
}

// Returns a new set containing the values in the first set but not the second.
//
// Does not modify either set.
func (s Set[V]) Difference(other Set[V]) Set[V] {
	difference := NewSet[V](s.Count())

	for v := range s {
		if !other.Contains(v) {
			difference.Insert(v)
		}
	}

	return difference
}

// Returns a new set containing the values in exactly one of the sets.
//
// Does not modify either set.
func (s Set[V]) SymmetricDifference(other Set[V]) Set[V] {
	union := s.Union(other)
	intersection := s.Intersection(other)

	return union.Difference(intersection)
}

// Returns true if the first set is a subset of the second.
func (s Set[V]) SubsetOf(other Set[V]) bool {
	if s.Count() > other.Count() {
		return false
	}

	for v := range s {
		if !other.Contains(v) {
			return false
		}
	}

	return true
}

// Returns true if the first set is a superset of the second.
func (s Set[V]) SupersetOf(other Set[V]) bool {
	return other.SubsetOf(s)
}

// Returns a JSON array of the set's values.
//
// Uses Set.Values.
func (s Set[V]) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.Values())
}

// Unmarshal a JSON array into a Set.
//
// Uses NewSetFromSlice, but will return an error if
// the slice length doesn't match the set count.
//
// Pointer receiver is used despite being inconsistent
// since it simplifies the use of json.Unmarshal.
// I.e. you can just use var set Set[int], whereas
// a value receiver would require manually calling NewSet.
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
