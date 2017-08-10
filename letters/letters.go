package letters

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"sort"
	"strings"
	"time"

	"github.com/tscott0/countdown/combi"
)

const (
	minWordLen int = 3
	maxWordLen int = 9
	maxAnswers int = 10
)

type success struct {
	Words []string `json:"words"`
	Took  string   `json:"took"`
}

type failure struct {
	Error string `json:"error"`
}

// dict is indexed by the word length then by word
// this allows searching by longest word first
var dict = make(map[string][]string)

// ReadDict reads the dictionary file and populates the dictionary map
func ReadDict(filename string) error {
	t0 := time.Now()

	file, err := os.Open(filename)
	if err != nil {
		return err
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
		if len(line) >= minWordLen && len(line) <= maxWordLen {
			if words, ok := dict[hashed]; ok {
				// If the hash already exists then append the word
				dict[hashed] = append(words, line)
			} else {
				// Insert it otherwise
				dict[hashed] = []string{line}
			}
		}
	}

	t1 := time.Now()
	fmt.Printf("Processed dictionary file %q in %v\n", filename, t1.Sub(t0))

	return nil
}

// Solve will find all possible words that can be formed from a string
// of letters without replacement.
func Solve(l string) ([]string, time.Duration, error) {
	var top []string
	var duration time.Duration

	t0 := time.Now()

	l = strings.ToUpper(l)
	letters := strings.Split(l, "")

	if len(letters) > maxWordLen {
		return top, duration,
			fmt.Errorf("%v is too long. The maxiumum number of letters is %v",
				l, maxWordLen)
	}

	if len(letters) < minWordLen {
		return top, duration,
			fmt.Errorf("%v is too short. The minimum number of letters is %v",
				l, minWordLen)
	}

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
		if c, found := dict[h]; found {
			for _, perm := range c {
				a.Insert(perm)
			}
		}
	}

	// Sort by word length so we can take the "top" answers
	sort.Sort(a)

	top = a.Top(maxAnswers)

	if len(top) == 0 {
		return top, duration,
			fmt.Errorf("No words found for letters %v",
				l)
	}

	t1 := time.Now()
	duration = t1.Sub(t0)

	fmt.Printf("Solved %q in %v. Top %v: %v\n", l, duration, maxAnswers, top)

	return top, duration, nil
}

// SolveJSON produces a JSON string containing
func SolveJSON(l string) (string, error) {
	w, d, err := Solve(l)

	// error response
	if err != nil {
		e := failure{err.Error()}

		b, err := json.Marshal(e)
		if err != nil {
			return "", fmt.Errorf("failed to marshal JSON: %v", err)
		}

		return string(b), nil
	}

	// success response
	s := success{
		w,
		d.String(),
	}

	b, err := json.MarshalIndent(s, "", "   ")
	if err != nil {
		return "", fmt.Errorf("failed to marshal JSON: %v", err)
	}

	return string(b), nil
}

// Sort the characters in the word alphabetically so that we
// can store unique combinations of letters
func hashWord(w string) string {
	sorted := strings.Split(w, "")
	sort.Strings(sorted)
	return strings.Join(sorted, "")
}
