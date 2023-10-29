/*
 * Copyright 2023 Christian Sigmon <cws@glasket.com>
 *
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/.
 */

package benchmarks_test

// import (
// 	"fmt"
// 	"math/rand"
// 	"runtime"
// 	"sync"
// 	"sync/atomic"
// 	"testing"

// 	"github.com/glasket/datastructures/interfaces/enumerator"
// 	"github.com/glasket/datastructures/list/arraylist"
// )

// func BenchmarkMutexBool(b *testing.B) {
// 	for _, size := range sizes {
// 		b.Run(fmt.Sprintf("%d", size), func(b *testing.B) {
// 			benchmarkMutexBool(b, size)
// 		})
// 	}
// }

// func benchmarkMutexBool(b *testing.B, size int) {
// 	l := enumerator.GetSliceEnumerable(rand.Perm(size))
// 	pe := enumerator.AsParallel[int](l)
// 	mut := sync.Mutex{}
// 	b.ResetTimer()
// 	for i := 0; i < b.N; i++ {
// 		t := rand.Intn(5000)
// 		quit := make(chan bool, runtime.NumCPU())
// 		var wg sync.WaitGroup
// 		success := true

// 		for _, partition := range pe.partitions {
// 			wg.Add(1)
// 			go func(p []int, success *bool, mut *sync.Mutex) {
// 				defer wg.Done()
// 				for _, v := range p {
// 					select {
// 					case <-quit:
// 						return
// 					default:
// 						if !(v > t) {
// 							quit <- true
// 							mut.Lock()
// 							*success = false
// 							mut.Unlock()
// 							return
// 						}
// 					}
// 				}
// 			}(partition, &success, &mut)
// 		}
// 		wg.Wait()
// 		if !success {
// 			_ = success
// 		}
// 	}
// }

// func BenchmarkAtomicBool(b *testing.B) {
// 	for _, size := range sizes {
// 		b.Run(fmt.Sprintf("%d", size), func(b *testing.B) {
// 			benchmarkAtomicBool(b, size)
// 		})
// 	}
// }

// func benchmarkAtomicBool(b *testing.B, size int) {
// 	l := enumerator.GetSliceEnumerable(rand.Perm(size))
// 	pe := enumerator.AsParallel[int](l)
// 	b.ResetTimer()
// 	for i := 0; i < b.N; i++ {
// 		t := rand.Intn(5000)
// 		quit := make(chan bool, runtime.NumCPU())
// 		var wg sync.WaitGroup
// 		failed := atomic.Bool{} // false
// 		for _, partition := range pe.Partitions() {
// 			wg.Add(1)
// 			go func(p []int, failed *atomic.Bool) {
// 				defer wg.Done()
// 				for _, v := range p {
// 					select {
// 					case <-quit:
// 						return
// 					default:
// 						if !(v > t) {
// 							quit <- true
// 							failed.Store(true)
// 							return
// 						}
// 					}
// 				}
// 			}(partition, &failed)
// 		}
// 		wg.Wait()
// 		if failed.Load() {
// 			_ = failed.Load()
// 		}
// 	}
// }

// var total int

// func isEven(v int) bool {
// 	return v%2 == 0
// }

var sizes = []int{100, 1000, 10000, 100000, 1000000}

// func BenchmarkSyncCount(b *testing.B) {
// 	for _, size := range sizes {
// 		b.Run(fmt.Sprintf("%d", size), func(b *testing.B) {
// 			benchmarkSyncCount(b, size)
// 		})
// 	}
// }

// func benchmarkSyncCount(b *testing.B, size int) {
// 	l := arraylist.NewFromSlice(rand.Perm(size))
// 	b.ResetTimer()
// 	for i := 0; i < b.N; i++ {
// 		total = enumerator.Count[int](l, isEven)
// 		_ = total
// 	}

// }

// func BenchmarkAtomicCount(b *testing.B) {
// 	for _, size := range sizes {
// 		b.Run(fmt.Sprintf("%d", size), func(b *testing.B) {
// 			benchmarkAtomicCount(b, size)
// 		})
// 	}
// }

// func benchmarkAtomicCount(b *testing.B, size int) {
// 	l := arraylist.NewFromSlice(rand.Perm(size))
// 	pe := enumerator.AsParallel(l.GetEnumerator())
// 	b.ResetTimer()
// 	for i := 0; i < b.N; i++ {
// 		count := atomic.Int64{}
// 		var wg sync.WaitGroup
// 		for _, partition := range pe.Partitions() {
// 			wg.Add(1)
// 			go func(p []int, count atomic.Int64) {
// 				defer wg.Done()
// 				for _, v := range p {
// 					if isEven(v) {
// 						count.Add(1)
// 					}
// 				}
// 			}(partition, count)
// 		}
// 		wg.Wait()
// 		total = int(count.Load())
// 		_ = total
// 	}
// }

// func BenchmarkArrayCount(b *testing.B) {
// 	for _, size := range sizes {
// 		b.Run(fmt.Sprintf("%d", size), func(b *testing.B) {
// 			benchmarkArrayCount(b, size)
// 		})
// 	}
// }

// func benchmarkArrayCount(b *testing.B, size int) {
// 	l := arraylist.NewFromSlice(rand.Perm(size))
// 	pe := enumerator.AsParallel(l.GetEnumerator())
// 	b.ResetTimer()
// 	for i := 0; i < b.N; i++ {
// 		count := make([]int, runtime.NumCPU())
// 		var wg sync.WaitGroup
// 		for idx, partition := range pe.Partitions() {
// 			wg.Add(1)
// 			go func(p []int, count *int) {
// 				defer wg.Done()
// 				for _, v := range p {
// 					if isEven(v) {
// 						*count += 1
// 					}
// 				}
// 			}(partition, &count[idx])
// 		}
// 		wg.Wait()
// 		total = enumerator.Sum(enumerator.GetSliceEnumerable(count))
// 		_ = total
// 	}
// }

// func BenchmarkDelayAtomicCount(b *testing.B) {
// 	for _, size := range sizes {
// 		b.Run(fmt.Sprintf("%d", size), func(b *testing.B) {
// 			benchmarkDelayAtomicCount(b, size)
// 		})
// 	}
// }

// func benchmarkDelayAtomicCount(b *testing.B, size int) {
// 	l := arraylist.NewFromSlice(rand.Perm(size))
// 	pe := enumerator.AsParallel(l.GetEnumerator())
// 	b.ResetTimer()
// 	for i := 0; i < b.N; i++ {
// 		count := atomic.Int64{}
// 		var wg sync.WaitGroup
// 		for _, partition := range pe.Partitions() {
// 			wg.Add(1)
// 			go func(p []int, count *atomic.Int64) {
// 				defer wg.Done()
// 				_count := 0
// 				for _, v := range p {
// 					if isEven(v) {
// 						_count += 1
// 					}
// 				}
// 				count.Add(int64(_count))
// 			}(partition, &count)
// 		}
// 		wg.Wait()
// 		total = int(count.Load())
// 		_ = total
// 	}
// }
