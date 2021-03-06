package words

import (
	"bufio"
	"os"
)

const wordlistValidGuessesFile = "./config/wordlist-valid-guesses.csv"
const wordlistValidSolutionsFile = "./config/wordlist-valid-solutions.csv"

func GetGuessesWordlist() (*Wordlist, error) {
	return makeWordListFromFile(wordlistValidGuessesFile)
}

func GetSolutionsWordlist() (*Wordlist, error) {
	return makeWordListFromFile(wordlistValidSolutionsFile)
}

func makeWordListFromFile(path string) (*Wordlist, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []Word
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		stringFromFile := scanner.Text()
		thisWord, _ := MakeWord(stringFromFile)
		lines = append(lines, *thisWord)
	}
	wordlist := makeWordlist(lines)
	return wordlist, scanner.Err()
}
