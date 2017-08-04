package perms

import (
	"fmt"

	"github.com/tscott0/countdown/combi"
)

func Permutations(n int) [][]int {
	t := [][]int{}

	for _, c := range combi.Combinations(n) {
		heapPermutation(&t, c, len(c))
	}

	return t
}

func heapPermutation(perms *[][]int, w []int, size int) {

	if size == 1 {
		// make sure we copy the underlying array
		new := make([]int, len(w))
		copy(new, w)
		*perms = append(*perms, new)
		//fmt.Printf("Calling heap with %v, %v. Current perms: %v\n", w, size, *perms)
	}

	for i := 0; i < size; i++ {
		heapPermutation(perms, w, size-1)

		if size%2 == 0 {
			w[i], w[size-1] = w[size-1], w[i]
		} else {
			w[0], w[size-1] = w[size-1], w[0]
		}
	}
}

// Permutations with repeats
func PermRep(n int, values []string) [][]string {
	pn := make([]int, n)
	p := make([]string, n)
	r := make([][]string, 0)
	for {
		// generate permutaton
		for i, x := range pn {
			p[i] = values[x]
		}

		new := make([]string, len(p))
		copy(new, p)
		r = append(r, new)

		// increment permutation number
		for i := 0; ; {
			pn[i]++
			if pn[i] < len(values) {
				break
			}
			pn[i] = 0
			i++
			if i == n {
				return r
			}
		}
	}
}

// Combinations with repeats
func Combrep(n int, lst []string) [][]string {
	fmt.Printf("Combrep(n=%v, lst=%v)\n", n, lst)
	if n == 0 {
		return [][]string{nil}
	}
	if len(lst) == 0 {
		return nil
	}
	r := Combrep(n, lst[1:])
	fmt.Println("MIDDLE")
	for _, x := range Combrep(n-1, lst) {
		r = append(r, append(x, lst[0]))
	}
	return r
}
