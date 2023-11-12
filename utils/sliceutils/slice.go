/*
 * Copyright 2023 Christian Sigmon <cws@glasket.com>
 *
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/.
 */

package sliceutils

import (
	"github.com/glasket/datastructures/utils/mathutils"
)

// Chunk splits a slice into a slice of slices.
// The number of slices in the returned slice will be chunkCount.
// The length of each slice in the returned slice is len(s) / chunkCount.
//
// If len(s) is not evenly divisible by chunkCount, the last slice will be smaller.
func Chunk[S ~[]E, E any](s S, chunkCount int) []S {
	chunkSize := mathutils.IntDivCeil(len(s), chunkCount)
	chunks := make([]S, 0, chunkCount)
	for i := 0; i < len(s); i += chunkSize {
		end := i + chunkSize
		if end > len(s) {
			end = len(s)
		}
		chunks = append(chunks, s[i:end])
	}
	return chunks
}

// Join joins a slice of slices into a single slice.
//
// If length is not provided, the starting capacity of the joined slice will be
// the sum of the lengths of the slices in s. This is calculated by iterating
// over s.
// If any lengths are provided, the starting capacity of the joined slice will
// be length[0].
//
// Any other values in the length array are ignored.
func Join[S ~[]E, E any](s []S, length ...int) S {
	var l int = 0
	if len(length) == 0 {
		for _, chunk := range s {
			l += len(chunk)
		}
	} else if length[0] > 0 {
		l = length[0]
	}

	joined := make(S, 0, l)
	for _, chunk := range s {
		joined = append(joined, chunk...)
	}
	return joined
}
