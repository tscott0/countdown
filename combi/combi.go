package combi

import (
	"fmt"
	"math"
)

// Combinations returns an array of int arrays that represent all combinations
// from a sample of size n.
//
// The number of possible combinations is (2^n - 1) where n is the size of the sample.
// For 3 elements this would mean 7 possible combinations of any length (except 0)
//
// Returns an array of array indices representing the combinations e.g.
//    [[0] [1] [0 1] [2] [0 2] [1 2] [0 1 2]]
func Combinations(n int) [][]int {
	output := [][]int{}

	maxNum := int(math.Pow(2, float64(n)))

	// Each number from 1 to maxNum represents a combination.
	// Represented in binary, each digit can either be 1 or 0.
	// A 1 indicates that the element is in the combination.
	for i := 1; i < maxNum; i++ {
		currentCombi := []int{}

		// Iterate over each bit starting with the least significant bit
		for j := 0; j < n; j++ {
			// A power of 2 is always has a single 1 when represented in binary
			power := int(math.Pow(2, float64(j)))

			// Debug string with numbers represented in a zero-padded binary format
			_ = fmt.Sprintf("%09b & %09b = %09b decimal=%v shifted=%v\n",
				i, power, i&power, i&power, i&power>>uint(j))

			// Do a bitwise AND with the power of 2 to determine the
			// value of the current bit.
			// If the output is non-zero then the digit was a 1
			if i&power != 0 {
				currentCombi = append(currentCombi, j)
			}
		}

		output = append(output, currentCombi)
	}

	return output
}
