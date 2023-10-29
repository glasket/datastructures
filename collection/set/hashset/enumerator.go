/*
 * Copyright 2023 Christian Sigmon <cws@glasket.com>
 *
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/.
 */

package hashset

import "github.com/glasket/datastructures/interfaces/enumerator"

// TODO Might be better to extract this out to the Set package
// Returns an enumerator.Enumerator for the set.
func (s *Set[V]) GetEnumerator() enumerator.IEnumerator[V] {
	return enumerator.GetMapKeyEnumerable(s.set).GetEnumerator()
}
