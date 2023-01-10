/*
 * Copyright 2023, Christian Sigmon <cws@glasket.com>
 *
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/.
 */
package set_test

import (
	"testing"

	. "github.com/glasket/datastructures/set"
)

var set Set[int]

func init() {
	set = *NewSet[int]()
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
	var found bool
	for _, v := range set.Values() {
		found = false
		for _, e := range expect {
			if v == e {
				found = true
			}
		}
		if !found {
			t.Errorf("Failed to find %v", v)
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
