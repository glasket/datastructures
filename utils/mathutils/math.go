/*
 * Copyright 2023 Christian Sigmon <cws@glasket.com>
 *
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/.
 */
package mathutils

import "math"

// IntDivCeil divides two integers and rounds the result up to the nearest whole number.
func IntDivCeil(x, y int) int {
	return int(math.Ceil(float64(x) / float64(y)))
}
