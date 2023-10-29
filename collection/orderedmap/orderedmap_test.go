/*
 * Copyright 2022, 2023 Christian Sigmon <cws@glasket.com>
 *
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/.
 */

package orderedmap_test

import (
	"reflect"
	"testing"

	. "github.com/glasket/datastructures/collection/orderedmap"
)

// TODO Cleanup this mess

// Tests that creation doesn't fail and has a sane default
func TestNewOrderedMap(t *testing.T) {
	om := NewOrderedMap[string, int](0)
	if om.Count() != 0 {
		t.Error("Count should be 0 at creation")
	}
}

// Tests basic operations for correctness
func TestOrderedMapOperations(t *testing.T) {
	om := NewOrderedMap[string, int](0)
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
	om.Set("key", expected+1)
	if v, _ := om.Get("key"); v != expected+1 {
		t.Error("Changing assigned value")
	}
	// Make sure key isn't reinserted a second time
	if len(om.Keys()) != om.Count() {
		t.Error("Key slice and map length mismatch after reassigning key")
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
	om := NewOrderedMap[int, int](0)
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

	// Test order correctness after reassigning key value
	om.Set(tests[0].key, 2)
	if om.Keys()[0] != tests[0].key {
		t.Errorf("First key should be %v, instead got %v", tests[0].key, om.Keys()[0])
	}

	// Test order correctness after SetAndUpdate
	om.SetAndUpdate(tests[0].key, 3)
	exp := []int{tests[1].key, tests[2].key, tests[3].key, tests[0].key}
	if !reflect.DeepEqual(om.Keys(), exp) {
		t.Errorf("Key order is broken.\nExpected: %v\nActual: %v", exp, om.Keys())
	}
}

// Tests the error conditions
func TestOrderedMapErrors(t *testing.T) {
	om := NewOrderedMap[string, int](0)
	if _, e := om.Get("key"); e == nil {
		t.Error("OrderedMap.Get should error when retrieving non-existent key")
	}
	om.Remove("key")
	if e := om.Remove("key"); e == nil {
		t.Error("OrderedMap.Remove should error when removing a non-existent key")
	}
}
