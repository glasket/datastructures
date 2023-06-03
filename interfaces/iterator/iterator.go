/*
 * Copyright 2023 Christian Sigmon <cws@glasket.com>
 *
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/.
 */
package iterator

// Iterable provides a way to iterate over an implementors values using range.
//
// Iterable does not guarantee thread safety, and generally should not be used
// when mutating the implementor.
type Iterable[V any] interface {
	Iter() <-chan V
}
