package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	var term = flag.Uint("term", 1, "the initial term")
	flag.Parse()

	if *term == 0 {
		fmt.Println("term cannot be zero")
		os.Exit(1)
	}

	var results []uint
	results = process(*term, []uint{})
	fmt.Println(results)
}

// process takes a term and results and recursively calculates
// new terms until the number 1 is reached
// does not validate the term argument -- assumes it is valid
func process(term uint, results []uint) []uint {
	newResults := append(results, term)

	// cater for when 1 is the term
	if term == 1 {
		return newResults
	}

	newTerm := calculateNewTerm(term)
	if newTerm != 1 {
		return process(newTerm, newResults)
	}

	return append(newResults, newTerm)
}

// calculateNewTerm takes a term and calculates if it is an odd or even number
// - if it's odd it will triple it and add 1
// - if it's even it will divided by 2
func calculateNewTerm(term uint) uint {
	var mod = term % 2
	if mod == 0 {
		return term / 2
	}
	return (term * 3) + 1
}
