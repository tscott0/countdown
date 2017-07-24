package main

import "strings"

// TODO: Move to main and pass this in to heapPermutation by reference instead
var perms = []string{}

func heapPermutation(w []string, size int) {
	if size == 1 {
		perms = append(perms, strings.Join(w, ""))
	}

	for i := 0; i < size; i++ {
		heapPermutation(w, size-1)

		if size%2 == 0 {
			w[i], w[size-1] = w[size-1], w[i]
		} else {
			w[0], w[size-1] = w[size-1], w[0]
		}
	}
}
