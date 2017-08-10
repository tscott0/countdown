package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"runtime/pprof"
	"strings"

	"github.com/gorilla/mux"
	"github.com/tscott0/countdown/letters"
	"github.com/tscott0/countdown/numbers"
)

const (
	dictionaryFile string = "words-en-gb"
)

func main() {

	f, err := os.Create("test.prof")
	if err != nil {
		log.Fatal(err)
	}
	pprof.StartCPUProfile(f)
	defer pprof.StopCPUProfile()

	if err := letters.ReadDict(dictionaryFile); err != nil {
		log.Fatalf("couldn't read dictionary file %q: %v", dictionaryFile, err)
	}

	router := newRouter()

	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatal(err)
	}
}

func lettersHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	l := vars["letters"]

	resp, err := letters.SolveJSON(l)
	if err != nil {
		log.Fatalf("failed to solve %q: %v", l, err)
	}

	w.Header().Add("Content-Type", "application/json")

	fmt.Fprintln(w, resp)
}

func numbersHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	n := vars["numbers"]

	a := strings.Split(n, ",")

	resp, err := numbers.SolveJSON(a)
	if err != nil {
		log.Fatalf("failed to solve %q: %v", n, err)
	}

	w.Header().Add("Content-Type", "application/json")

	fmt.Fprintln(w, resp)
}
