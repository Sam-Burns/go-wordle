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
	fmt.Println("Frequencies of characters in solution wordlist:")
	fmt.Println()

	frequencyTable := words.MakeFrequencyTableFromWordlist(*validSolutionsWordlist)

	frequencyTableRows := frequencyTable.GetPrintableFrequencyTableRows()

	for _, frequencyTableRow := range frequencyTableRows {
		fmt.Println(string(frequencyTableRow.Character) + "\t" + frequencyTableRow.GetPrintablePercentage())
	}

	fmt.Println()
}

func dieIfError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}
