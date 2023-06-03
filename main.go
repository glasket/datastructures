package main

import "golang.org/x/exp/rand"

func main() {
	x := make([]int, 0)
	for {
		x = append(x, rand.Perm(1024)...)
	}
}
