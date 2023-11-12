package sliceutils_test

import (
	"testing"

	"github.com/glasket/datastructures/utils/sliceutils"
)

func TestChunk(t *testing.T) {
	s := []int{1, 2, 3, 4, 5, 6, 7, 8}

	results := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8},
	}

	chunks := sliceutils.Chunk(s, 3)
	for i := range chunks {
		for j := range chunks[i] {
			if chunks[i][j] != results[i][j] {
				t.Error("Chunk did not properly chunk the slice")
			}
		}
		if len(chunks[i]) != len(results[i]) {
			t.Error("Chunk did not properly chunk the slice")
		}
	}

}

func TestJoin(t *testing.T) {
	s := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8},
	}

	results := []int{1, 2, 3, 4, 5, 6, 7, 8}

	joined := sliceutils.Join(s)
	for i := range joined {
		if joined[i] != results[i] {
			t.Error("Join did not properly join the slices")
		}
	}
	if len(joined) != len(results) {
		t.Error("Join did not properly join the slices")
	}
}
