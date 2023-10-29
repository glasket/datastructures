/*
 * Copyright 2023 Christian Sigmon <cws@glasket.com>
 *
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/.
 */

package enumerator

import (
	"runtime"
	"sync"
	"sync/atomic"

	"github.com/glasket/datastructures/utils/sliceutils"
)

var cc int = runtime.NumCPU()

// SetConcurrency sets the number of goroutines to use for parallel operations.
// This will impact the number of slice chunks for parallel operations.
//
// Defaults to runtime.NumCPU().
func SetConcurrency(c int) {
	cc = c
}

type parallelEnumerator[V any] struct {
	partitions [][]V
	length     int
}

func asParallel[V any](e IEnumerable[V]) *parallelEnumerator[V] {
	vals := e.ToSlice()
	return &parallelEnumerator[V]{
		partitions: sliceutils.Chunk(vals, cc),
		length:     len(vals),
	}
}

// ParallelEach calls f on each element of e in parallel.
func ParallelEach[V any](e IEnumerable[V], f func(V)) {
	pe := asParallel(e)
	var wg sync.WaitGroup
	for _, partition := range pe.partitions {
		wg.Add(1)
		go func(p []V) {
			defer wg.Done()
			for _, v := range p {
				f(v)
			}
		}(partition)
	}
	wg.Wait()
}

// ParallelAll returns true if f returns true for all elements of e.
//
// ParallelAll short circuits on the first false return from f.
func ParallelAll[V any](e IEnumerable[V], f func(V) bool) bool {
	pe := asParallel(e)
	quit := make(chan bool, cc)
	defer close(quit)
	var wg sync.WaitGroup
	success := atomic.Bool{}
	success.Store(true)

	for _, partition := range pe.partitions {
		wg.Add(1)
		go func(p []V, success *atomic.Bool) {
			defer wg.Done()
			for _, v := range p {
				select {
				case <-quit:
					return
				default:
					if !f(v) {
						quit <- true
						success.CompareAndSwap(true, false)
						return
					}
				}
			}
		}(partition, &success)
	}

	wg.Wait()
	return success.Load()
}

// ParallelAny returns true if f returns true for any element of e.
//
// ParallelAny short circuits on the first true return from f.
func ParallelAny[V any](e IEnumerable[V], f func(V) bool) bool {
	pe := asParallel(e)
	quit := make(chan bool, cc)
	defer close(quit)
	var wg sync.WaitGroup
	success := atomic.Bool{}

	for _, partition := range pe.partitions {
		wg.Add(1)
		go func(p []V, success *atomic.Bool) {
			defer wg.Done()
			for _, v := range p {
				select {
				case <-quit:
					return
				default:
					if f(v) {
						quit <- true
						success.CompareAndSwap(false, true)
						return
					}
				}
			}
		}(partition, &success)
	}

	wg.Wait()
	return success.Load()
}

// ParallelCount returns the number of elements of e for which f returns true.
func ParallelCount[V any](e IEnumerable[V], f func(V) bool) int {
	pe := asParallel(e)
	count := atomic.Int64{}
	var wg sync.WaitGroup

	for _, partition := range pe.partitions {
		wg.Add(1)
		go func(p []V, count *atomic.Int64) {
			defer wg.Done()
			_count := 0
			for _, v := range p {
				if f(v) {
					_count += 1
				}
			}
			count.Add(int64(_count))
		}(partition, &count)
	}

	wg.Wait()
	return int(count.Load())
}

// ParallelMap returns a new IEnumerable with the results of applying f to each element of e.
//
// ParallelMap preserves order.
func ParallelMap[V any, R any](e IEnumerable[V], f func(V) R) IEnumerable[R] {
	pe := asParallel(e)
	result := make([]R, pe.length)
	var wg sync.WaitGroup

	chunkedRes := sliceutils.Chunk(result, cc)
	for idx, partition := range pe.partitions {
		wg.Add(1)
		go func(p []V, r []R) {
			defer wg.Done()
			for i, v := range p {
				r[i] = f(v)
			}
		}(partition, chunkedRes[idx])
	}
	wg.Wait()
	return GetSliceEnumerable(result)
}

// ParallelFilter returns a new IEnumerable with the elements of e for which f returns true.
//
// ParallelFilter preserves order.
func ParallelFilter[V any](e IEnumerable[V], f func(V) bool) IEnumerable[V] {
	pe := asParallel(e)
	var wg sync.WaitGroup
	chunkedRes := make([][]V, cc)

	for idx, partition := range pe.partitions {
		wg.Add(1)
		go func(p []V, r *[]V) {
			defer wg.Done()
			for _, v := range p {
				if f(v) {
					*r = append(*r, v)
				}
			}
		}(partition, &chunkedRes[idx])
	}
	wg.Wait()
	return GetSliceEnumerable(sliceutils.Join(chunkedRes))
}

// ParallelReduce returns a single value by applying the given function f to each element in the enumerable.
// The combiner function is used to combine the results of each partition of the enumerable.
func ParallelReduce[V any, R any](e IEnumerable[V], f func(R, V) R, combiner func([]R) R, initial R) R {
	pe := asParallel(e)
	var wg sync.WaitGroup
	chunkedRes := make([]R, cc)

	for idx, partition := range pe.partitions {
		wg.Add(1)
		go func(p []V, idx int) {
			defer wg.Done()
			r := &chunkedRes[idx]
			for i := range p {
				*r = f(*r, p[i])
			}
		}(partition, idx)
	}
	wg.Wait()
	return combiner(chunkedRes)
}
