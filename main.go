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

	for i := MinWordLen; i <= MaxWordLen; i++ {
		for _ = range combi.Combinations(i) {

		}
	}

	testWord := "POST"

	g, err := guess(testWord, dict)
	if err != nil {
		fmt.Println(err)
	}

	if g {
		fmt.Printf("%v found\n", testWord)
	}

	//	for _, w := range perms {
	//		fmt.Println(w)
	//	}

}

func guess(guess string, dict map[int]map[string][]string) (bool, error) {
	hashedGuess := hashWord(guess)

	fmt.Printf("Searching for %s using %s\n", guess, hashedGuess)

	// Check if the hash of the word exists first
	if words, ok := dict[len(hashedGuess)][hashedGuess]; ok {
		fmt.Printf("Hash of %s found in %v\n", guess, dict[len(hashedGuess)][hashedGuess])
		// Then iterate over valid anagrams of the hash
		for _, w := range words {
			if w == guess {
				return true, nil
			}
		}
	}
	return false, nil
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
