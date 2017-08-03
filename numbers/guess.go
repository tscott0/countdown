package numbers

import (
	"fmt"
	"strconv"
)

type guess struct {
	operands  []int
	operators []string
}

func (g *guess) newGuess(i int) {
	g.operands = []int{i}
	g.operators = []string{}
}

func (g *guess) total() int {
	if len(g.operands) < 1 {
		return 0
	}

	total := g.operands[0]

	for i, operator := range g.operators {
		operand := g.operands[i+1]

		switch operator {
		case "+":
			total += operand
		case "-":
			total -= operand
		case "*":
			total *= operand
		case "/":
			total /= operand
		}
	}

	return total
}

func (g *guess) string() string {
	if len(g.operands) < 1 {
		return ""
	}

	text := strconv.Itoa(g.operands[0])

	for i, operator := range g.operators {
		operand := g.operands[i+1]

		switch operator {
		case "+":
			text = "(" + text + operator + strconv.Itoa(operand) + ")"
		case "-":
			text = "(" + text + operator + strconv.Itoa(operand) + ")"
		case "*":
			text = "(" + text + operator + strconv.Itoa(operand) + ")"
		case "/":
			text = "(" + text + operator + strconv.Itoa(operand) + ")"
		}
	}

	return text
}

// Returns false if a division returns a non-integer
func (g *guess) insert(i int, op string) bool {
	// TODO: check that op is a valid operator

	g.operands = append(g.operands, i)
	g.operators = append(g.operators, op)

	return true
	// TODO: division check
}

func (g *guess) bestGuess(newGuess guess, target int) {
	// Replace a nil pointer the first time
	if g == nil {
		*g = newGuess
		return
	}

	fmt.Printf("%v = %v vs %v = %v\n",
		g.string(), g.total(),
		newGuess.string(), newGuess.total())

	newToTarget := target - newGuess.total()
	if newToTarget < 0 {
		newToTarget = -newToTarget
	}

	currentToTarget := target - g.total()
	if currentToTarget < 0 {
		currentToTarget = -currentToTarget
	}

	if newToTarget < currentToTarget {
		*g = newGuess
	}
}
