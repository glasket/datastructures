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
	"testing"

	. "github.com/glasket/datastructures/set"
)

var set Set[int]

func init() {
	set = NewSet[int](0)
}

func TestNewSetFromSlice(t *testing.T) {
	vals := []int{1, 2, 3, 4}
	s := NewSetFromSlice(vals)
	if s.Count() != len(vals) {
		t.Errorf("expected %v items, found %v", len(vals), s.Count())
	}
	for _, v := range vals {
		if !s.Contains(v) {
			t.Errorf("%v not in set", v)
		}
	}
}

func TestInsert(t *testing.T) {
	vals := []int{1, 2, 3, 4}
	for _, v := range vals {
		set.Insert(v)
		if !set.Contains(v) {
			t.Errorf("%v was not inserted", v)
		}
	}
	if err := set.Insert(1); err == nil {
		t.Error("set failed to catch that value is already present before insertion")
	}
}

func TestJSON(t *testing.T) {
	j, err := json.Marshal(set)
	if err != nil {
		t.Error(err)
	}
	if string(j) != "[1,2,3,4]" {
		t.Error("json marshalled incorrectly")
	}

	var s Set[int]
	err = json.Unmarshal(j, &s)
	if err != nil {
		t.Error(err)
	}
	missing := findValues(s.Values(), set.Values())
	if len(missing) != 0 {
		t.Error("json unmarshalled incorrectly")
		for _, v := range missing {
			t.Errorf("%v missing", v)
		}
	}
}

func TestRemove(t *testing.T) {
	set.Remove(1)
	if set.Contains(1) {
		t.Error("remove failed to remove 1")
	}
	if err := set.Remove(1); err == nil {
		t.Error("set failed to catch that value wasn't present before removal")
	}
}

func TestValues(t *testing.T) {
	expect := []int{2, 3, 4}
	if len(set.Values()) != len(expect) {
		t.Errorf("Not enough values. %v/%v", len(set.Values()), len(expect))
	}
	missing := findValues(set.Values(), expect)
	if len(missing) != 0 {
		for _, v := range missing {
			t.Errorf("%v missing", v)
		}
	}
}

func TestClear(t *testing.T) {
	set.Clear()
	if set.Count() != 0 {
		t.Errorf("set should be empty, found %v elements", set.Count())
		for _, v := range set.Values() {
			t.Log(v)
		}
	}
}

func findValues[V comparable](setVals []V, expectedVals []V) []V {
	missing := []V{}
	found := false
	for _, v := range setVals {
		found = false
		for _, e := range expectedVals {
			if v == e {
				found = true
			}
		}
		if !found {
			missing = append(missing, v)
		}
	}
	return missing
}
