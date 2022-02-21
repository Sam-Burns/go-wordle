package main

import (
	"fmt"
	"github.com/Sam-Burns/go-wordle/pkg/wordlegame"
	"github.com/Sam-Burns/go-wordle/pkg/wordleplayer"
	"github.com/Sam-Burns/go-wordle/pkg/words"
	"os"
	"strconv"
)

func main() {

	targetWord := getTargetWordFromParam()

	wordleGame := wordlegame.MakeGameFromTargetWord(*targetWord)

	player := wordleplayer.MakePlayer()

	turnNumber := 1
	won := false

	for turnNumber <= 6 && !won {
		printPreAnalysis(player)
		guessWasSolution, guess, feedback, evaluation := takeGuess(turnNumber, player, wordleGame, targetWord)
		printEvaluation(evaluation)
		won = guessWasSolution
		printTurn(guess, feedback, turnNumber)
		turnNumber += 1
	}

	printOutcome(won, turnNumber-1)
}

func getTargetWordFromParam() *words.Word {
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
	fmt.Println()

	return targetWord
}

func takeGuess(guessNo int, player *wordleplayer.WordlePlayer, wordleGame *wordlegame.Game, targetWord *words.Word) (won bool, guess *words.Word, feedback *wordlegame.Feedback, evaluation *wordleplayer.ProposedGuessEvaluation) {
	guess, evaluation = player.GetNextGuess(*wordleGame, guessNo == 6)
	won = guess.Equals(targetWord)
	feedback = wordlegame.GetFeedback(targetWord, guess)
	player.TakeFeedbackFromGuess(*guess, *feedback)
	return
}

func printTurn(guess *words.Word, feedback *wordlegame.Feedback, guessNo int) {
	fmt.Println("Guess number " + strconv.Itoa(guessNo) + ": " + guess.ToString())
	fmt.Println("Feedback from guess was: " + feedback.String())
	fmt.Println()
}

func printPreAnalysis(player *wordleplayer.WordlePlayer) {
	noOfPossibleSolutions := player.GetNoOfPossibleSolutions()
	fmt.Print("There are currently " + strconv.Itoa(noOfPossibleSolutions) + " possible solutions")
	if noOfPossibleSolutions <= 6 {
		fmt.Println(" [" + player.GetPossibleSolutions() + "]")
	} else {
		fmt.Println()
	}
}

func printEvaluation(evaluation *wordleplayer.ProposedGuessEvaluation) {
	if !evaluation.IsNullEvaluation() {
		fmt.Println("The next guess should be " + evaluation.ProposedGuess.String)
		fmt.Println("Worst-case scenario for proposed guess is the feedback " + evaluation.GetWorstCaseScenarioFeedbackString() + ". Carry-over ratio for possible solutions list would be " + evaluation.GetWorstCaseShortlistCarryOverRatioString())
	}
}

func printOutcome(won bool, turnNumber int) {
	if won {
		fmt.Println("Won the Wordle in " + strconv.Itoa(turnNumber) + " turns")
	} else {
		fmt.Println("Lost the Wordle after 6 turns :-(")
	}
	fmt.Println()
}
