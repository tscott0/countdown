package perms

func Permutations(n int) [][]int {
	t := [][]int{}

	var data []int
	for i := 0; i < n; i++ {
		data = append(data, i)
	}

	for j := 1; j <= len(data); j++ {
		heapPermutation(&t, data, j)
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

func Combrep(n int, lst []string) [][]string {
	if n == 0 {
		return [][]string{nil}
	}
	if len(lst) == 0 {
		return nil
	}
	r := Combrep(n, lst[1:])
	for _, x := range Combrep(n-1, lst) {
		r = append(r, append(x, lst[0]))
	}
	return r
}
