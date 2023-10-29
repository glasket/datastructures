/*
 * Copyright 2023 Christian Sigmon <cws@glasket.com>
 *
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/.
 */

package enumerator

import "golang.org/x/exp/constraints"

// Each calls the given function for each element in the enumerable.
func Each[V any](e IEnumerable[V], f func(V)) {
	enum := e.GetEnumerator()
	for enum.Next() {
		f(enum.Current())
	}
}

// All returns true if the given function returns true for all elements in the enumerable.
func All[V any](e IEnumerable[V], f func(V) bool) bool {
	enum := e.GetEnumerator()
	for enum.Next() {
		if !f(enum.Current()) {
			return false
		}
	}
	return true
}

// Any returns true if the given function returns true for any element in the enumerable.
func Any[V any](e IEnumerable[V], f func(V) bool) bool {
	enum := e.GetEnumerator()
	for enum.Next() {
		if f(enum.Current()) {
			return true
		}
	}
	return false
}

// Count returns the number of elements in the enumerable for which the given function returns true.
func Count[V any](e IEnumerable[V], f func(V) bool) int {
	enum := e.GetEnumerator()
	count := 0
	for enum.Next() {
		if f(enum.Current()) {
			count += 1
		}
	}
	return count
}

// Map returns a new enumerable with the given function applied to each element in the enumerable.
func Map[V any, R any](e IEnumerable[V], f func(V) R) IEnumerable[R] {
	enum := e.GetEnumerator()
	results := make([]R, 0)
	for enum.Next() {
		results = append(results, f(enum.Current()))
	}
	return GetSliceEnumerable(results)
}

// Filter returns a new enumerable with only the elements for which the given function returns true.
func Filter[V any](e IEnumerable[V], f func(V) bool) IEnumerable[V] {
	enum := e.GetEnumerator()
	results := make([]V, 0)
	for enum.Next() {
		if f(enum.Current()) {
			results = append(results, enum.Current())
		}
	}
	return GetSliceEnumerable(results)
}

// Reduce returns a single value by applying the given function to each element in the enumerable.
func Reduce[V any, R any](e IEnumerable[V], f func(R, V) R, initial R) R {
	enum := e.GetEnumerator()
	result := initial
	for enum.Next() {
		result = f(result, enum.Current())
	}
	return result
}

type number interface {
	constraints.Integer | constraints.Float | constraints.Complex
}

func add[V constraints.Integer](x, y V) V {
	return x + y
}

func sub[V constraints.Integer](x, y V) V {
	return x - y
}

// Range returns an IEnumerable of integers on the interval [start, end).
// If end < start, the range will be descending.
func Range[V constraints.Integer](start, end V) IEnumerable[V] {
	// Func pointers allow for unsigned integer ranges,
	// whereas i += (1/-1) does not due to unsigned casting problems
	var f func(V, V) V
	var cap V

	if end < start {
		f = sub[V]
		cap = start - end
	} else if end > start {
		f = add[V]
		cap = end - start
	} else {
		return GetSliceEnumerable([]V{start})
	}

	out := make([]V, 0, cap)
	// Guaranteed ±1 increments, so != is perfectly fine
	for i := start; i != end; i = f(i, 1) {
		out = append(out, i)
	}
	return GetSliceEnumerable(out)
}

// Sum returns the sum of all elements in the enumerable.
func Sum[V number](e IEnumerable[V]) V {
	enum := e.GetEnumerator()
	var result V
	for enum.Next() {
		result += enum.Current()
	}
	return result
}
