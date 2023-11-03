/*
 * Copyright 2023 Christian Sigmon <cws@glasket.com>
 *
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/.
 */

package collection

import "github.com/glasket/datastructures/interfaces/enumerator"

type IImmutableCollection[V comparable] interface {
	enumerator.IEnumerable[V]
	Contains(v V) bool
	Count() int
	IsEmpty() bool
	String() string
}

type ICollection[V comparable] interface {
	IImmutableCollection[V]
	Add(v V)
	Remove(v V)
	Clear()
}

type IImmutableIndexedCollection[V comparable] interface {
	IImmutableCollection[V]
	IndexOf(v V) (int, error)
	Get(i int) (V, error)
}

type IIndexedCollection[V comparable] interface {
	ICollection[V]
	IImmutableIndexedCollection[V]
	Set(i int, v V) error
	InsertAt(i int, v V) error
	RemoveAt(i int) error
}
