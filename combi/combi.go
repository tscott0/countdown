package combi

import (
	"fmt"
	"math"
)

// Combinations returns an array of int arrays that represent all combinations
// of indices from an arry of size n
func Combinations(n int) [][]int {

	//fmt.Printf("%v\n", math.Pow(2, float64(size)))

	output := [][]int{}
	maxNum := int(math.Pow(2, float64(n)) - 1)

	for i := 1; i < maxNum; i++ {
		currentCombi := []int{}

		for j := 0; j < n; j++ {
			power := int(math.Pow(2, float64(j)))
			fmt.Printf("%08b & %08b = %08b decimal=%v shifted=%v\n", i, power, i&power, i&power, i&power>>uint(j))
			if i&power>>uint(j) == 1 {
				currentCombi = append(currentCombi, j)
			}
		}
		output = append(output, currentCombi)
	}

	return output
}
