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
func Chunk[V any](s []V, chunkCount int) [][]V {
	chunkSize := mathutils.IntDivCeil(len(s), chunkCount)
	chunks := make([][]V, 0, chunkCount)
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
func Join[V any](s [][]V, length ...int) []V {
	var l int = 0
	if len(length) == 0 {
		for _, chunk := range s {
			l += len(chunk)
		}
	} else if length[0] > 0 {
		l = length[0]
	}

	joined := make([]V, 0, l)
	for _, chunk := range s {
		joined = append(joined, chunk...)
	}
	return joined
}
