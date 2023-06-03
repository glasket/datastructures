/*
 * Copyright 2023 Christian Sigmon <cws@glasket.com>
 *
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/.
 */
package set

import "github.com/glasket/datastructures/collection"

type ISet[V comparable] interface {
	collection.ICollection[V]
	Equals(other ISet[V]) bool
	Union(other ISet[V]) ISet[V]
	Intersection(other ISet[V]) ISet[V]
	Complement(other ISet[V]) ISet[V]
	RelativeComplement(other ISet[V]) ISet[V]
	SymmetricDifference(other ISet[V]) ISet[V]
	SubsetOf(other ISet[V]) bool
	SupersetOf(other ISet[V]) bool
}

type Tuple[V comparable, T comparable] [2]any
