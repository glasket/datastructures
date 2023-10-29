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

func TestEach(t *testing.T) {
	s := enumerator.GetSliceEnumerable(rand.Perm(PERM_SIZE))
	out_arr := make([]int, 1024)
	enumerator.Each(s, func(i int) {
		out_arr[i] = i
	})
	expected := enumerator.Range(0, 1024).GetEnumerator()
	if len(out_arr) != PERM_SIZE {
		t.Error("Each did not iterate over all elements")
	}
	for expected.Next() {
		if out_arr[expected.Current()] != expected.Current() {
			t.Error("Each did not properly assign elements")
		}
	}
}

func TestAll(t *testing.T) {
	s := enumerator.GetSliceEnumerable(rand.Perm(PERM_SIZE))
	if !enumerator.All(s, func(i int) bool {
		return i < PERM_SIZE
	}) {
		t.Error("All did not return true for i < PERM_SIZE")
	}
	if enumerator.All(s, func(i int) bool {
		return i > PERM_SIZE
	}) {
		t.Error("All did not return false for i > PERM_SIZE")
	}
}

func TestAny(t *testing.T) {
	s := enumerator.GetSliceEnumerable(rand.Perm(PERM_SIZE))
	randGoal := rand.Intn(PERM_SIZE)
	if !enumerator.Any(s, func(i int) bool {
		return i == randGoal
	}) {
		t.Errorf("Any did not return true for i == %d", randGoal)
	}
	if enumerator.Any(s, func(i int) bool {
		return i > PERM_SIZE
	}) {
		t.Error("Any did not return false for i > PERM_SIZE")
	}
}

func TestCount(t *testing.T) {
	s := enumerator.GetSliceEnumerable(rand.Perm(PERM_SIZE))
	if enumerator.Count(s, func(i int) bool {
		return i < PERM_SIZE
	}) != PERM_SIZE {
		t.Error("Count did not return correct number of elements")
	}
	if enumerator.Count(s, func(i int) bool {
		return i > PERM_SIZE
	}) != 0 {
		t.Error("Count did not return correct number of elements")
	}
}

func TestMap(t *testing.T) {
	s := enumerator.Range(0, PERM_SIZE)
	mapped := enumerator.Map(s, func(i int) int {
		return i * i
	})
	mappedEnum := mapped.GetEnumerator()
	expected := enumerator.Range(0, PERM_SIZE).GetEnumerator()
	for expected.Next() {
		ok := mappedEnum.Next()
		if !ok {
			t.Error("Map did not map all elements")
		}
		if mappedEnum.Current() != expected.Current()*expected.Current() {
			t.Error("Map did not properly map elements")
		}
	}
}

func TestFilter(t *testing.T) {
	s := enumerator.GetSliceEnumerable(rand.Perm(PERM_SIZE))
	filtered := enumerator.Filter(s, func(i int) bool {
		return i%2 == 0
	})
	if len(filtered.ToSlice()) != PERM_SIZE/2 {
		t.Errorf("Filter did not filter out all odd elements, expected %d elements, got %d", PERM_SIZE/2, len(filtered.ToSlice()))
	}
	for _, v := range filtered.ToSlice() {
		if v%2 != 0 {
			t.Errorf("Filter did not filter out all odd elements, found %d", v)
		}
	}
}

func TestReduce(t *testing.T) {
	s := enumerator.GetSliceEnumerable(rand.Perm(PERM_SIZE))
	reduced := enumerator.Reduce(s, func(a, b int) int {
		return a + b
	}, 0)
	if reduced != (PERM_SIZE*(PERM_SIZE-1))/2 {
		t.Errorf("Reduce did not reduce properly, expected %d, got %d", (PERM_SIZE*(PERM_SIZE-1))/2, reduced)
	}
}

func TestRange(t *testing.T) {
	// Test ascending range
	s := enumerator.Range(0, 10)
	if len(s.ToSlice()) != 10 {
		t.Errorf("Range did not return %d elements, got %d", 10, len(s.ToSlice()))
	}
	for i := 0; i < 10; i++ {
		if s.ToSlice()[i] != i {
			t.Errorf("Range did not return correct elements, expected %d, got %d", i, s.ToSlice()[i])
		}
	}

	// Test descending range
	s = enumerator.Range(10, 0)
	if len(s.ToSlice()) != 10 {
		t.Errorf("Range did not return %d elements, got %d", 10, len(s.ToSlice()))
	}
	for i := 10; i > 0; i-- {
		if s.ToSlice()[10-i] != i {
			t.Errorf("Range did not return correct elements, expected %d, got %d", i, s.ToSlice()[10-i])
		}
	}

	// Test single value range
	s = enumerator.Range(10, 10)
	if len(s.ToSlice()) != 1 {
		t.Errorf("Range did not return %d elements, got %d", 1, len(s.ToSlice()))
	}
	if s.ToSlice()[0] != 10 {
		t.Errorf("Range did not return correct elements, expected %d, got %d", 10, s.ToSlice()[0])
	}
}

func TestSum(t *testing.T) {
	s := enumerator.GetSliceEnumerable(rand.Perm(PERM_SIZE))
	sum := enumerator.Sum(s)
	expected := PERM_SIZE * (PERM_SIZE - 1) / 2
	if sum != expected {
		t.Errorf("Sum did not sum properly, expected %d, got %d", expected, sum)
	}
}
