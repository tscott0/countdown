package letters

import (
	"sort"
	"strings"

	"github.com/tscott0/countdown/combi"
)

const (
	minWordLen int = 3
	maxWordLen int = 9
	maxAnswers int = 10
)

// Solve will find all possible words that can be formed from a string
// of letters without replacement.
func Solve(l string, dict map[int]map[string][]string) []string {
	letters := strings.Split(strings.ToUpper(l), "")
	guesses := []string{}

	for _, c := range combi.Combinations(len(letters)) {
		currentGuess := ""

		for _, x := range c {
			currentGuess = currentGuess + letters[x]
		}

		if len(currentGuess) >= minWordLen {
			guesses = append(guesses, currentGuess)
		}
	}

	a := answers{
		[]string{},
		make(map[string]struct{}),
	}

	for _, g := range guesses {
		h := hashWord(g)
		if c, found := dict[len(h)][h]; found {
			for _, perm := range c {
				a.Insert(perm)
			}
		}
	}

	// Sort by word length so we can take the best answers
	sort.Sort(a)

	return a.TopWords()
}

func hashWord(w string) string {
	sorted := strings.Split(w, "")
	sort.Strings(sorted)
	return strings.Join(sorted, "")
}
