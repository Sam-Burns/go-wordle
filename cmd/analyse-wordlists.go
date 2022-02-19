package main

import (
	"fmt"
	"github.com/Sam-Burns/go-wordle/pkg/words"
	"os"
)

func main() {

	validGuessesWordlist, err := words.GetGuessesWordlist()
	dieIfError(err)
	validSolutionsWordlist, err := words.GetSolutionsWordlist()
	dieIfError(err)

	fmt.Println()
	fmt.Printf("There are %d words that are valid guesses\n", validGuessesWordlist.Count())
	fmt.Printf("There are %d words that are valid solutions\n", validSolutionsWordlist.Count())
	fmt.Println()

}

func dieIfError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}
