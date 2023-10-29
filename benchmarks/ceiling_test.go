/*
 * Copyright 2023 Christian Sigmon <cws@glasket.com>
 *
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/.
 */

package benchmarks

import (
	"math"
	"math/rand"
	"testing"
	"time"
)

var r *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))

const MAX_RAND int = 4096

func BenchmarkCeiling(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = ceiling(r.Intn(MAX_RAND)+1, r.Intn(MAX_RAND)+1)
	}
}

func ceiling(x, y int) int {
	return int(math.Ceil(float64(x) / float64(y)))
}

func BenchmarkIntCeiling(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = intCeiling(r.Intn(MAX_RAND)+1, r.Intn(MAX_RAND)+1)
	}
}

func intCeiling(x, y int) int {
	return 1 + (x-1)/y
}
