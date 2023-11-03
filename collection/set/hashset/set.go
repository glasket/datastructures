/*
 * Copyright 2023 Christian Sigmon <cws@glasket.com>
 *
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/.
 */

package hashset

import (
	"encoding/json"
	"fmt"

	"github.com/glasket/datastructures/collection/set"
)

var _ set.ISet[int] = (*Set[int])(nil)

var fil struct{} = struct{}{}

type Set[V comparable] struct {
	set     map[V]struct{}
	values  []V
	version int
}

func New[V comparable](size int) *Set[V] {
	return &Set[V]{
		set:     make(map[V]struct{}, size),
		values:  nil,
		version: 0,
	}
}

// NewFromSlice creates a set from a preexisting slice
//
// Ignores the Add error for repeated values.
// If you want to catch errors, either use NewSet and manually Add,
// or check the size of the slice against the resulting set.
func NewFromSlice[V comparable](slice []V) *Set[V] {
	set := New[V](len(slice))
	for _, v := range slice {
		set.Add(v) // We don't care if it errors
	}
	return set
}

// Add adds a given value to the set.
//
// no-op if value is already present.
func (s *Set[V]) Add(value V) {
	if s.Contains(value) {
		return
	}
	s.set[value] = fil
	s.version += 1
	s.values = nil
}

// Remove removes the given value from the set.
//
// no-op if value is not present
func (s *Set[V]) Remove(value V) {
	if !s.Contains(value) {
		return
	}
	delete(s.set, value)
	s.version += 1
	s.values = nil
}

// Clear deletes all values from the set.
func (s *Set[V]) Clear() {
	s.set = make(map[V]struct{})
	s.version += 1
	s.values = nil
}

// Contains returns true if the given value is present in the set.
func (s *Set[V]) Contains(value V) bool {
	_, ok := s.set[value]
	return ok
}

// String eturns the string representation of the set.
func (s *Set[V]) String() string {
	return fmt.Sprintf("Set[%v]", s.Values())
}

// IsEmpty eturns true if the set is empty.
func (s *Set[V]) IsEmpty() bool {
	return s.Count() == 0
}

// Count returns the number of values in the set.
func (s *Set[V]) Count() int {
	return len(s.set)
}

// ToSlice eturns a slice of all values in the set.
func (s *Set[V]) Values() []V {
	if s.values != nil {
		return s.values
	}
	s.values = make([]V, 0, s.Count())
	for v := range s.set {
		s.values = append(s.values, v)
	}
	return s.values
}

// Equals tests if two sets are equal.
//
// Two sets are equal if they contain the same values.
func (s *Set[V]) Equals(other set.ISet[V]) bool {
	if s.Count() != other.Count() {
		return false
	}

	for v := range s.set {
		if !other.Contains(v) {
			return false
		}
	}
	return true
}

// [Union] returns a new set containing the values in both sets.
//
// Does not modify either set.
//
// [Union]: https://en.wikipedia.org/wiki/Union_(set_theory)
func (s *Set[V]) Union(other set.ISet[V]) set.ISet[V] {
	union := New[V](s.Count() + other.Count())

	for v := range s.set {
		union.Add(v)
	}
	for _, v := range other.Values() {
		union.Add(v)
	}

	return union
}

// [Intersection] returns a new set containing the values in both sets.
//
// Does not modify either set.
//
// [Intersection]: https://en.wikipedia.org/wiki/Intersection_(set_theory)
func (s *Set[V]) Intersection(other set.ISet[V]) set.ISet[V] {
	intersection := New[V](s.Count())

	for v := range s.set {
		if other.Contains(v) {
			intersection.Add(v)
		}
	}

	return intersection
}

// [Complement] returns a new set containing the values in the first set but not the second.
//
// Does not modify either set.
//
// [Complement]: https://en.wikipedia.org/wiki/Complement_(set_theory)
func (s *Set[V]) Complement(other set.ISet[V]) set.ISet[V] {
	complement := New[V](s.Count())

	for v := range s.set {
		if !other.Contains(v) {
			complement.Add(v)
		}
	}

	return complement
}

// [RelativeComplement] returns a new set containing the values in the second set but not the first
//
// Does not modify either set.
//
// [RelativeComplement]: https://en.wikipedia.org/wiki/Complement_(set_theory)#Relative_complement
func (s *Set[V]) RelativeComplement(other set.ISet[V]) set.ISet[V] {
	return other.Complement(s)
}

// [SymmetricDifference] returns a new set containing the values in exactly one of the sets.
//
// Does not modify either set.
//
// [SymmetricDifference]: https://en.wikipedia.org/wiki/Symmetric_difference
func (s *Set[V]) SymmetricDifference(other set.ISet[V]) set.ISet[V] {
	union := s.Union(other)
	intersection := s.Intersection(other)

	return union.Complement(intersection)
}

// SubsetOf returns true if the first set is a subset of the second.
func (s *Set[V]) SubsetOf(other set.ISet[V]) bool {
	if s.Count() > other.Count() {
		return false
	}

	for v := range s.set {
		if !other.Contains(v) {
			return false
		}
	}

	return true
}

// SupersetOf returns true if the first set is a superset of the second.
func (s *Set[V]) SupersetOf(other set.ISet[V]) bool {
	return other.SubsetOf(s)
}

// MarshalJSON returns a JSON array of the set's values.
//
// Uses Set.Values.
func (s *Set[V]) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.Values())
}

// UnmarshalJSON turns a JSON array into a Set.
//
// Uses NewFromSlice, but will return an error if
// the slice length doesn't match the set count.
func (s *Set[V]) UnmarshalJSON(data []byte) error {
	var slice []V
	err := json.Unmarshal(data, &slice)
	if err != nil {
		return err
	}
	*s = *NewFromSlice(slice)
	if s.Count() != len(slice) {
		return fmt.Errorf("duplicates were present")
	}
	return nil
}
