/*
 * Copyright 2023 Christian Sigmon <cws@glasket.com>
 *
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/.
 */

package enumerator_test

import (
	"math/rand"
	"testing"

	"github.com/glasket/datastructures/interfaces/enumerator"
)

func TestParallelEach(t *testing.T) {
	s := enumerator.GetSliceEnumerable(rand.Perm(PERM_SIZE))
	out_arr := make([]int, 1024)
	enumerator.ParallelEach[int](s, func(i int) {
		out_arr[i] = i
	})
	expected := enumerator.Range(0, 1024).GetEnumerator()
	if len(out_arr) != PERM_SIZE {
		t.Error("ParallelEach did not iterate over all elements")
	}
	for expected.Next() {
		if out_arr[expected.Current()] != expected.Current() {
			t.Error("ParallelEach did not properly assign elements")
		}
	}
}

func TestParallelAll(t *testing.T) {
	s := enumerator.GetSliceEnumerable(rand.Perm(PERM_SIZE))
	if !enumerator.ParallelAll[int](s, func(i int) bool {
		return i < PERM_SIZE
	}) {
		t.Error("Expected ParallelAll to return true")
	}

	if enumerator.ParallelAll[int](s, func(i int) bool {
		return i < 5
	}) {
		t.Error("Expected ParallelAll to return false")
	}
}

func TestParallelAny(t *testing.T) {
	s := enumerator.GetSliceEnumerable(rand.Perm(PERM_SIZE))

	if !enumerator.ParallelAny[int](s, func(i int) bool {
		return i == 5
	}) {
		t.Error("Expected ParallelAny to return true")
	}

	if enumerator.ParallelAny[int](s, func(i int) bool {
		return i == PERM_SIZE
	}) {
		t.Error("Expected ParallelAny to return false")
	}
}

// A function to test the parallel count function
func TestParallelCount(t *testing.T) {
	s := enumerator.GetSliceEnumerable(rand.Perm(PERM_SIZE))

	out := enumerator.ParallelCount[int](s, func(i int) bool {
		return i%2 == 0
	})
	if out != PERM_SIZE/2 {
		t.Errorf("Expected ParallelCount to return %d, got %d", PERM_SIZE/2, out)
	}
}

func TestParallelMap(t *testing.T) {
	s := enumerator.GetSliceEnumerable(rand.Perm(PERM_SIZE))

	out := enumerator.ParallelMap[int](s, func(i int) int {
		return i * 2
	})
	if len(out.Values()) != PERM_SIZE {
		t.Fatalf("Expected ParallelMap to return %d elements, got %d", PERM_SIZE, len(out.Values()))
	}
	for i := 0; i < PERM_SIZE; i++ {
		expected := s.Values()[i] * 2
		got := out.Values()[i]
		if expected != got {
			t.Errorf("Expected ParallelMap to return %d at index %d, got %d", expected, i, got)
		}
	}
}

func TestParallelFilter(t *testing.T) {
	s := enumerator.GetSliceEnumerable(rand.Perm(PERM_SIZE))

	out := enumerator.ParallelFilter[int](s, func(i int) bool {
		return i%2 == 0
	})
	if len(out.Values()) != PERM_SIZE/2 {
		t.Fatalf("Expected ParallelFilter to return %d elements, got %d", PERM_SIZE/2, len(out.Values()))
	}
	for i := 0; i < PERM_SIZE/2; i++ {
		got := out.Values()[i]
		if got%2 != 0 {
			t.Errorf("Expected ParallelFilter to only return evens, got %d at index %d", got, i)
		}
	}
}

func TestParallelReduce(t *testing.T) {
	s := enumerator.GetSliceEnumerable(rand.Perm(PERM_SIZE))

	out := enumerator.ParallelReduce[int](s, func(acc int, i int) int {
		return acc + i
	}, func(r []int) int {
		sum := 0
		for _, v := range r {
			sum += v
		}
		return sum
	}, 0)
	expected := PERM_SIZE * (PERM_SIZE - 1) / 2
	if out != expected {
		t.Errorf("Expected ParallelReduce to return %d, got %d", expected, out)
	}
}
