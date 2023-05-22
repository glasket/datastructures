/*
 * Copyright 2023, Christian Sigmon <cws@glasket.com>
 *
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/.
 */
package set_test

import (
	"encoding/json"
	"sort"
	"testing"

	. "github.com/glasket/datastructures/set"
)

func TestNewSet(t *testing.T) {
	set := NewSet[int](5)
	if len(set) != 0 {
		t.Fatalf("Expected NewSet to return a set of size 0, got %d", len(set))
	}
}

func TestNewSetFromSlice(t *testing.T) {
	set := NewSetFromSlice([]int{1, 2, 3, 4, 5})
	if len(set) != 5 {
		t.Fatalf("Expected NewSetFromSlice to return a set of size 5, got %d", len(set))
	}
	for i := 1; i <= 5; i++ {
		if !set.Contains(i) {
			t.Errorf("Expected NewSetFromSlice to contain %d", i)
		}
	}
}

func TestSetInsert(t *testing.T) {
	set := NewSet[int](0)
	err := set.Insert(5)
	if err != nil {
		t.Errorf("Expected Insert to not error, got %v", err)
	}
	if len(set) != 1 {
		t.Errorf("Expected Insert to insert a value, got %d", len(set))
	}
}

func TestSetInsertDuplicate(t *testing.T) {
	set := NewSet[int](0)
	err := set.Insert(5)
	if err != nil {
		t.Errorf("Expected Insert to not error, got %v", err)
	}
	err = set.Insert(5)
	if err == nil {
		t.Errorf("Expected Insert to error on duplicate, got nil")
	}
	if len(set) != 1 {
		t.Errorf("Expected Insert to not insert a duplicate value, got %d", len(set))
	}
}

func TestSetRemove(t *testing.T) {
	set := NewSet[int](0)
	set.Insert(5)
	err := set.Remove(5)
	if err != nil {
		t.Errorf("Expected Remove to not error, got %v", err)
	}
	if len(set) != 0 {
		t.Errorf("Expected Remove to remove a value, got %d", len(set))
	}
}

func TestSetRemoveMissing(t *testing.T) {
	set := NewSet[int](0)
	err := set.Remove(5)
	if err == nil {
		t.Errorf("Expected Remove to error on missing value, got nil")
	}
	if len(set) != 0 {
		t.Errorf("Expected Remove to not remove a missing value, got %d", len(set))
	}
}

func TestSetClear(t *testing.T) {
	set := NewSet[int](0)
	set.Insert(5)
	set.Clear()
	if len(set) != 0 {
		t.Errorf("Expected Clear to clear set, got %d", len(set))
	}
}

func TestSetContains(t *testing.T) {
	set := NewSet[int](0)
	set.Insert(5)
	if !set.Contains(5) {
		t.Errorf("Expected Contains to return true, got false")
	}
	if set.Contains(6) {
		t.Errorf("Expected Contains to return false, got true")
	}
}

func TestSetCount(t *testing.T) {
	set := NewSet[int](0)
	set.Insert(5)
	if set.Count() != 1 {
		t.Errorf("Expected Count to return 1, got %d", set.Count())
	}
}

func TestSetValues(t *testing.T) {
	set := NewSet[int](0)
	set.Insert(5)
	set.Insert(6)
	values := set.Values()
	if len(values) != 2 {
		t.Errorf("Expected Values to return a slice of size 2, got size %d", len(values))
	}
	sort.Slice(values, func(i, j int) bool { return values[i] < values[j] })
	if values[0] != 5 || values[1] != 6 {
		t.Errorf("Expected Values to return [5 6], got %v", values)
	}
}

func TestSetEquals(t *testing.T) {
	set1 := NewSetFromSlice([]int{1, 2, 3})
	set2 := NewSetFromSlice([]int{1, 2, 3})

	if !set1.Equals(set2) {
		t.Errorf("Expected Equals to return true, got false")
	}

	set2.Insert(4)

	if set1.Equals(set2) {
		t.Errorf("Expected Equals to return false, got true")
	}
}

func TestSetUnion(t *testing.T) {
	set1 := NewSetFromSlice([]int{1, 2, 3})

	set2 := NewSetFromSlice([]int{3, 4, 5})

	union := set1.Union(set2)
	if len(union) != 5 {
		t.Errorf("Expected Union to return a set of size 5, got %d", len(union))
	}
	expected := NewSetFromSlice([]int{1, 2, 3, 4, 5})
	if !union.Equals(expected) {
		t.Errorf("Expected Union to return %v, got %v", expected, union)
	}
}

func TestSetIntersection(t *testing.T) {
	set1 := NewSetFromSlice([]int{1, 2, 3})

	set2 := NewSetFromSlice([]int{3, 4, 5})

	intersection := set1.Intersection(set2)
	if len(intersection) != 1 {
		t.Errorf("Expected Intersection to return a set of size 1, got %d", len(intersection))
	}
	expected := NewSetFromSlice([]int{3})
	if !intersection.Equals(expected) {
		t.Errorf("Expected Intersection to return %v, got %v", expected, intersection)
	}
}

func TestSetDifference(t *testing.T) {
	set1 := NewSetFromSlice([]int{1, 2, 3})

	set2 := NewSetFromSlice([]int{3, 4, 5})

	difference := set1.Difference(set2)
	if len(difference) != 2 {
		t.Errorf("Expected Difference to return a set of size 2, got %d", len(difference))
	}
	expected := NewSetFromSlice([]int{1, 2})
	if !difference.Equals(expected) {
		t.Errorf("Expected Difference to return %v, got %v", expected, difference)
	}
}

func TestSetSymmetricDifference(t *testing.T) {
	set1 := NewSetFromSlice([]int{1, 2, 3})

	set2 := NewSetFromSlice([]int{3, 4, 5})

	difference := set1.SymmetricDifference(set2)
	if len(difference) != 4 {
		t.Errorf("Expected SymmetricDifference to return a set of size 4, got %d", len(difference))
	}
	expected := NewSetFromSlice([]int{1, 2, 4, 5})
	if !difference.Equals(expected) {
		t.Errorf("Expected SymmetricDifference to return %v, got %v", expected, difference)
	}
}

func TestSetSubsetOf(t *testing.T) {
	set1 := NewSetFromSlice([]int{1, 2, 3})

	set2 := NewSetFromSlice([]int{1, 2, 3, 4, 5})

	if !set1.SubsetOf(set2) {
		t.Errorf("Expected SubsetOf to return true, got false")
	}
	if set2.SubsetOf(set1) {
		t.Errorf("Expected SubsetOf to return false, got true")
	}
}

func TestSetSupersetOf(t *testing.T) {
	set1 := NewSetFromSlice([]int{1, 2, 3})

	set2 := NewSetFromSlice([]int{1, 2, 3, 4, 5})

	if set1.SupersetOf(set2) {
		t.Errorf("Expected SupersetOf to return false, got true")
	}
	if !set2.SupersetOf(set1) {
		t.Errorf("Expected SupersetOf to return true, got false")
	}
}

func TestSetJson(t *testing.T) {
	set := NewSetFromSlice([]int{1, 2, 3, 4, 5})

	jsonBytes, err := json.Marshal(set)
	if err != nil {
		t.Errorf("Expected Marshal to not error, got %v", err)
	}

	var set2 Set[int]
	err = json.Unmarshal(jsonBytes, &set2)
	if err != nil {
		t.Errorf("Expected Unmarshal to not error, got %v", err)
	}

	if !set.Equals(set2) {
		t.Errorf("Expected Unmarshal to return %v, got %v", set, set2)
	}
}
