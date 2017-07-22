package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {

	// dict is indexed by the word length then by word
	// this allows searching by longest word first
	dict := make(map[int]map[string]struct{})

	dict[3] = make(map[string]struct{})
	var es struct{}
	dict[3]["OPT"] = es
	dict[3]["POT"] = es
	dict[3]["TOP"] = es

	//fmt.Println(hashWord("POT"))
	//fmt.Println(hashWord("TOP"))

	guess := "PTO"
	if _, ok := dict[len(guess)][guess]; ok {
		fmt.Printf("%s is a valid word\n", guess)
	}

}

func hashWord(w string) string {
	sorted := strings.Split(w, "")
	sort.Strings(sorted)
	return strings.Join(sorted, "")
}
