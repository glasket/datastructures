/*
 * Copyright 2022, Christian Sigmon <cws@glasket.com>
 *
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/.
 */
package datastructures_test

import (
	"testing"

	. "glasket.com/datastructures"
)

// Tests that creation doesn't fail and has a sane default
func TestNewOrderedMap(t *testing.T) {
	om := *NewOrderedMap[string, int](0)
	if om.Count() != 0 {
		t.Error("Count should be 0 at creation")
	}
}

// Tests basic operations for correctness
func TestOrderedMapOperations(t *testing.T) {
	om := *NewOrderedMap[string, int](0)
	expected := 1
	om.Set("key", expected)
	if v, e := om.Get("key"); v != expected || e != nil {
		t.Errorf("Value: %v | Expected: %v | Error: %v", v, expected, e)
	}
	if om.Count() != 1 {
		t.Error("Count is incorrect after insert")
	}
	if !om.Contains("key") {
		t.Error("Map incorrectly showing that key is not present")
	}
	if len(om.Keys()) != om.Count() {
		t.Error("Key slice and map length mismatch")
	}
	if err := om.Remove("key"); err != nil {
		t.Error("Removal should not throw an error on an existing key")
	}
	if om.Count() != 0 {
		t.Error("Count not properly reducing on removal")
	}
	if om.Contains("key") {
		t.Error("Map incorrectly showing that key is still present after removal")
	}
}

// Tests for insertion order and maintenance of order after removal
func TestOrderedMapOrdering(t *testing.T) {
	tests := []struct {
		key, val int
	}{
		{3, 1},
		{4, 1},
		{1, 1},
		{2, 1},
		{5, 1},
	}
	om := *NewOrderedMap[int, int](0)
	for _, v := range tests {
		om.Set(v.key, v.val)
	}
	// Test that count was maintained
	if om.Count() != len(tests) {
		t.Error("Count is inaccurate")
	}
	// Test that the map is correct
	for _, v := range tests {
		if !om.Contains(v.key) {
			t.Errorf("%v is missing", v.key)
		}
		if mapv, e := om.Get(v.key); mapv != v.val || e != nil {
			t.Errorf("Value: %v | Expected: %v | Error: %v", mapv, v.val, e)
		}
	}
	// Test insertion order
	for i, v := range om.Keys() {
		if v != tests[i].key {
			t.Errorf("Order violation: %v != %v at index %v", v, tests[i].key, i)
		}
	}
	// Test order correctness after removal
	om.Remove(1)
	tests = []struct {
		key, val int
	}{
		{3, 1},
		{4, 1},
		{2, 1},
		{5, 1},
	}
	for i, v := range om.Keys() {
		if v != tests[i].key {
			t.Errorf("Order violation: %v != %v at index %v", v, tests[i].key, i)
		}
	}
}

// Tests the error conditions
func TestOrderedMapErrors(t *testing.T) {
	om := *NewOrderedMap[string, int](0)
	if _, e := om.Get("key"); e == nil {
		t.Error("OrderedMap.Get should error when retrieving non-existent key")
	}
	om.Set("key", 1)
	if e := om.Set("key", 1); e == nil {
		t.Error("OrderedMap.Set should error when inserting a duplicate key")
	}
	om.Remove("key")
	if e := om.Remove("key"); e == nil {
		t.Error("OrderedMap.Remove should error when removing a non-existent key")
	}
}
