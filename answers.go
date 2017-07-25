package main

type answers []string

func (a answers) Len() int      { return len(a) }
func (a answers) Swap(i, j int) { a[i], a[j] = a[j], a[i] }

// Less first compares by string length and then alphabetically
func (a answers) Less(i, j int) bool {
	if len(a[i]) == len(a[j]) {
		return a[i] < a[j]
	}

	return len(a[i]) < len(a[j])
}
