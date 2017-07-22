package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {

	dict := make(map[int]map[string][]string)

	dict[3] = make(map[string][]string)
	test := []string{"POT", "TOP", "OPT"}
	dict[3]["OPT"] = test

	//fmt.Println(hashWord("POT"))
	//fmt.Println(hashWord("TOP"))

	guess := "POT"
	for _, word := range dict[len(guess)][hashWord(guess)] {
		if word == guess {
			fmt.Printf("%s is a valid word\n", guess)
			break
		}
	}
}

func hashWord(w string) string {
	sorted := strings.Split(w, "")
	sort.Strings(sorted)
	return strings.Join(sorted, "")
}
