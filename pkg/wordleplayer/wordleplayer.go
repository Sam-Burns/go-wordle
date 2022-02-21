package wordleplayer

import (
	"github.com/Sam-Burns/go-wordle/pkg/wordlegame"
	"github.com/Sam-Burns/go-wordle/pkg/words"
)

type WordlePlayer struct {
	possibleWords words.Wordlist
	validGuesses  words.Wordlist
}

func MakePlayer() *WordlePlayer {

	possibleWords, _ := words.GetGuessesWordlist()
	validGuesses, _ := words.GetGuessesWordlist()

	return &WordlePlayer{*possibleWords, *validGuesses}
}

func (player WordlePlayer) GetNextGuess(game wordlegame.Game) words.Word {
	word, _ := words.MakeWord("ORATE")
	return *word
}
