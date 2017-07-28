package combi_test

import (
	"fmt"

	"github.com/tscott0/countdown/combi"
)

func ExampleCombinations() {
	letters := []string{"A", "B", "C"}

	combinations := []string{}

	for _, c := range combi.Combinations(len(letters)) {
		combi := ""
		for _, x := range c {
			combi = combi + letters[x]
		}

		combinations = append(combinations, combi)
	}

	fmt.Println(combinations)
	// Output: [A B AB C AC BC ABC]
}
