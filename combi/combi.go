package combi

import (
	"fmt"
	"math"
)

// Combinations returns an array of uint arrays that represent all combinations
// of indices from an arry of size n
func Combinations(n uint) [][]uint {

	//fmt.Printf("%v\n", math.Pow(2, float64(size)))

	output := [][]uint{}
	maxNum := int(math.Pow(2, float64(n)) - 1)

	for i := 1; i < maxNum; i++ {
		currentCombi := []uint{}

		var j uint
		for j = 0; j < n; j++ {
			power := int(math.Pow(2, float64(j)))
			fmt.Printf("%08b & %08b = %08b decimal=%v shifted=%v\n", i, power, i&power, i&power, i&power>>j)
			if i&power>>j == 1 {
				currentCombi = append(currentCombi, j)
			}
		}
		output = append(output, currentCombi)
	}

	return output
}
