package wordlegame

import "github.com/Sam-Burns/go-wordle/pkg/words"

type Game struct {
	TargetWord words.Word
}

func MakeGameFromTargetWord(targetWord words.Word) *Game {
	return &Game{targetWord}
}
