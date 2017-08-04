package numbers

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/tscott0/countdown/perms"
)

var ops = []string{"+", "-", "*", "/"}

type success struct {
	Solved  bool   `json:"solved"`
	Formula string `json:"formula"`
	Total   int    `json:"total"`
	Score   int    `json:"score"`
	Took    string `json:"took"`
}

type failure struct {
	Error string `json:"error"`
}

func Solve(operands []int, target int) (Guess, time.Duration, error) {
	var closest Guess
	var duration time.Duration

	// check for too many operands
	if len(operands) > 6 {
		return closest, duration, fmt.Errorf("the maximum number of operands is 6")
	}

	// check for too many operands
	if len(operands) < 2 {
		return closest, duration, fmt.Errorf("the minimum number of operands is 2")
	}

	if target < 101 || target > 999 {
		return closest, duration, fmt.Errorf("the target must be between 101 and 999 inclusive")
	}

	t0 := time.Now()

	for _, c := range perms.Permutations(len(operands)) {
		var currentCombi []int

		for _, x := range c {
			currentCombi = append(currentCombi, operands[x])
		}

		if len(c) == 1 {
			continue
		}

		// there is always one fewer operator than operands
		numberOfOperators := len(currentCombi) - 1
		operatorPerms := perms.PermRep(numberOfOperators, ops)

		for _, opGroups := range operatorPerms {

			var g Guess
			g.newGuess(currentCombi[0])

			if len(opGroups) != len(currentCombi)-1 {
				return closest, duration, fmt.Errorf("operators and operands do not match")
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
				return closest, duration, nil
			}
		}
	}

	t1 := time.Now()
	duration = t1.Sub(t0)

	fmt.Printf("Closest after %v: %v = %v\n", duration, closest.string(), closest.total())
	return closest, duration, nil
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

	// Take first number as the target
	target := castedNumbers[0]
	castedNumbers = castedNumbers[1:]

	a, d, err := Solve(castedNumbers, target)

	// error response
	if err != nil {
		e := failure{err.Error()}

		b, err := json.Marshal(e)
		if err != nil {
			return "", fmt.Errorf("failed to marshal JSON: %v", err)
		}

		return string(b), nil
	}

	result := a.total()

	// success response
	s := success{
		target == result,
		a.string(),
		result,
		0,
		d.String(),
	}

	b, err := json.MarshalIndent(s, "", "   ")
	if err != nil {
		return "", fmt.Errorf("failed to marshal JSON: %v", err)
	}

	return string(b), nil
}
