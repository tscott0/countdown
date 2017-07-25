package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strings"

	"github.com/tscott0/countdown/combi"
)

const (
	MinWordLen int = 3
	MaxWordLen int = 9
	MaxAnswers int = 10
)

func main() {

	// dict is indexed by the word length then by word
	// this allows searching by longest word first
	dict := make(map[int]map[string][]string)

	// Make all the inner maps
	for i := MinWordLen; i <= MaxWordLen; i++ {
		dict[i] = make(map[string][]string)
	}

	buildDict("words-en-gb", dict)

	// TODO: Iterate over permutations/combinations of letters in reverse length order
	// heapPermutation(strings.Split(guess, ""), len(guess))

	letters := strings.Split("GYHDNOEUR", "")
	guesses := []string{}

	// TODO: Improve by iterating over longest words first
	for _, c := range combi.Combinations(len(letters)) {
		currentGuess := ""

		for _, x := range c {
			currentGuess = currentGuess + letters[x]
		}

		if len(currentGuess) >= MinWordLen {
			guesses = append(guesses, currentGuess)
		}
	}

	var a answers = []string{}

	for _, g := range guesses {
		h := hashWord(g)
		if _, found := dict[len(h)][h]; found {
			for _, perm := range dict[len(h)][h] {
				a = append(a, perm)
			}
		}
	}

	// Sort by word length so we can take the best answers
	sort.Sort(a)

	fmt.Println(a[len(a)-MaxAnswers:])

}

// Read the local dictionary file and populate the dictionary
func buildDict(filename string, dict map[int]map[string][]string) error {

	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	defer file.Close()

	reader := bufio.NewReader(file)

	for {
		lineBytes, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}

		line := strings.ToUpper(string(lineBytes))
		hashed := hashWord(line)

		// Only store words of valid length
		if len(line) >= MinWordLen && len(line) <= MaxWordLen {
			if words, ok := dict[len(hashed)][hashed]; ok {
				// If the hash already exists then append the word
				dict[len(hashed)][hashed] = append(words, line)
			} else {
				// Insert it otherwise
				dict[len(hashed)][hashed] = []string{line}
			}
		}
	}

	return nil
}

func hashWord(w string) string {
	sorted := strings.Split(w, "")
	sort.Strings(sorted)
	return strings.Join(sorted, "")
}
