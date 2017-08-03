package numbers

import (
	"fmt"
	"strconv"
	"time"

	"github.com/tscott0/countdown/combi"
	"github.com/tscott0/countdown/perms"
)

var ops = []string{"+", "-", "*", "/"}

type success struct {
	Words []string `json:"words"`
	Took  string   `json:"took"`
}

type failure struct {
	Error string `json:"error"`
}

func Solve(numbers []int) (string, time.Duration, error) {
	var duration time.Duration
	var closest guess

	// TODO: Handle invalid numbers length here

	// Take first number as the target
	target := numbers[0]
	numbers = numbers[1:]

	//fmt.Println(numbers)
	t0 := time.Now()

	// TODO: USE PERMS NOT COMBIS YOU IDIOT
	for _, c := range combi.Combinations(len(numbers)) {
		var currentCombi []int

		for _, x := range c {
			currentCombi = append(currentCombi, numbers[x])
		}

		if len(c) == 1 {
			continue
		}
		operatorPerms := perms.Combrep(len(currentCombi)-1, ops)

		for _, opGroups := range operatorPerms {

			var g guess
			g.newGuess(currentCombi[0])

			if len(opGroups) != len(currentCombi)-1 {
				return closest.string(), duration, fmt.Errorf("operators and operands do not match")
			}

			for i := 0; i < len(opGroups); i++ {
				g.insert(currentCombi[i+1], opGroups[i])
				//fmt.Printf("%v = %v\n", g.string(), g.total())
			}

			closest.bestGuess(g, target)

			// Return immediately if we hit the target
			if closest.total() == target {
				t1 := time.Now()
				duration = t1.Sub(t0)

				fmt.Printf("Solved after %v: %v = %v\n", duration, closest.string(), closest.total())
				return closest.string(), duration, nil
			}
		}
	}

	t1 := time.Now()
	duration = t1.Sub(t0)

	fmt.Printf("Closest after %v: %v = %v\n", duration, closest.string(), closest.total())
	return closest.string(), duration, nil
}

// SolveJSON produces a JSON string containing
func SolveJSON(numbers []string) (string, error) {

	var castedNumbers []int

	for _, n := range numbers {
		cast, err := strconv.ParseInt(n, 10, 64)

		if err != nil {
			// TODO: Return error JSON
			return "", err
		}
		castedNumbers = append(castedNumbers, int(cast))
	}

	a, d, err := Solve(castedNumbers)

	_, _ = d, err

	return a, nil
}
