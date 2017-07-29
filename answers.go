package main

type answers struct {
	words  []string
	unique map[string]struct{}
}

func (a answers) Len() int      { return len(a.words) }
func (a answers) Swap(i, j int) { a.words[i], a.words[j] = a.words[j], a.words[i] }

// Less first compares by string length and then alphabetically
func (a answers) Less(i, j int) bool {
	if len(a.words[i]) == len(a.words[j]) {
		return a.words[i] < a.words[j]
	}

	return len(a.words[i]) < len(a.words[j])
}

func (a *answers) Insert(w string) {
	// Add an answer to a map of unique answers and to the array for sorting
	if _, ok := a.unique[w]; !ok {
		var es struct{}
		a.unique[w] = es
		a.words = append(a.words, w)
	}
}

func (a answers) TopWords() []string {
	if len(a.words) < maxAnswers {
		return a.words
	}
	return a.words[len(a.words)-maxAnswers:]
}
