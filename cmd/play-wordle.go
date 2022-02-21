package main

import (
	"fmt"
	wordlegame "github.com/Sam-Burns/go-wordle/pkg/wordlegame"
	wordleplayer "github.com/Sam-Burns/go-wordle/pkg/wordleplayer"
	"github.com/Sam-Burns/go-wordle/pkg/words"
	"os"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "Usage: go run cmd/play-wordle.go ORATE")
		os.Exit(1)
	}

	wordArgument := os.Args[1]

	targetWord, err := words.MakeWord(wordArgument)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error processing target word: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Wordle solution: " + targetWord.ToString())

	wordleGame := wordlegame.MakeGameFromTargetWord(*targetWord)

	player := wordleplayer.MakePlayer()

	nextWord := player.GetNextGuess(*wordleGame)

	fmt.Println("Player's first guess: " + nextWord.ToString())
}
