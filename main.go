package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gorilla/mux"
)

const (
	dictionaryFile string = "words-en-gb"
)

// dict is indexed by the word length then by word
// this allows searching by longest word first
var dict = make(map[int]map[string][]string)

func main() {
	// Make all the inner maps
	for i := minWordLen; i <= maxWordLen; i++ {
		dict[i] = make(map[string][]string)
	}

	if err := readDict(dictionaryFile, dict); err != nil {
		log.Fatalf("couldn't read dictionary file %q: %v", dictionaryFile, err)
	}

	router := newRouter()

	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatal(err)
	}
}

// Read the local dictionary file and populate the dictionary
func readDict(filename string, dict map[int]map[string][]string) error {
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
			if words, ok := dict[len(hashed)][hashed]; ok {
				// If the hash already exists then append the word
				dict[len(hashed)][hashed] = append(words, line)
			} else {
				// Insert it otherwise
				dict[len(hashed)][hashed] = []string{line}
			}
		}
	}

	t1 := time.Now()
	fmt.Printf("Processed dictionary file %q in %v\n", filename, t1.Sub(t0))

	return nil
}

func lettersShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	letters := vars["letters"]
	fmt.Println(vars)
	fmt.Fprintln(w, "Solving: ", letters.Solve(letters, dict))
}
