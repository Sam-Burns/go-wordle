package wordleplayer

import (
	"github.com/Sam-Burns/go-wordle/pkg/wordlegame"
	"github.com/Sam-Burns/go-wordle/pkg/words"
	"strings"
)

type WordlePlayer struct {
	possibleSolutions words.Wordlist
	validGuesses      words.Wordlist
}

func MakePlayer() *WordlePlayer {

	possibleSolutions, _ := words.GetSolutionsWordlist()
	validGuesses, _ := words.GetGuessesWordlist()

	return &WordlePlayer{*possibleSolutions, *validGuesses}
}

func (player WordlePlayer) GetNextGuess(game wordlegame.Game, isSixthTurn bool) (guess *words.Word, evaluation *ProposedGuessEvaluation) {

	bestGuessEvaluation := ProposedGuessEvaluation{worstCaseShortlistCarryOverRatio: 1.0}

	if player.possibleSolutions.Count() == 1 {
		return &player.possibleSolutions.Words[0], &bestGuessEvaluation
	}

	if isSixthTurn {
		return &player.possibleSolutions.Words[0], &bestGuessEvaluation
	}

	for _, proposedGuess := range player.validGuesses.Words {
		proposedGuessEvaluation := player.EvaluatePossibleGuess(&proposedGuess)

		if proposedGuessEvaluation.getWorstCaseShortlistCarryOverRatio() < bestGuessEvaluation.getWorstCaseShortlistCarryOverRatio() {
			bestGuessEvaluation = *proposedGuessEvaluation
		}
	}

	return &bestGuessEvaluation.ProposedGuess, &bestGuessEvaluation
}

func (player WordlePlayer) EvaluatePossibleGuess(possibleGuess *words.Word) *ProposedGuessEvaluation {

	proposedGuessEvaluation := MakeProposedGuessEvaluation(*possibleGuess, player.possibleSolutions.Count())

	for _, possibleSolution := range player.possibleSolutions.Words {
		feedback := wordlegame.GetFeedback(&possibleSolution, possibleGuess)
		proposedGuessEvaluation.AddPossibleOutcome(possibleSolution, *feedback)
	}

	return &proposedGuessEvaluation
}

func (player *WordlePlayer) TakeFeedbackFromGuess(word words.Word, feedback wordlegame.Feedback) {

	var newShortlist []words.Word

	for _, solutionStillOnShortlist := range player.possibleSolutions.Words {
		feedbackIfThisWordWereSolution := wordlegame.GetFeedback(&solutionStillOnShortlist, &word)
		if feedbackIfThisWordWereSolution.Equals(feedback) {
			newShortlist = append(newShortlist, solutionStillOnShortlist)
		}
	}

	player.possibleSolutions = words.Wordlist{Words: newShortlist}
}

func (player *WordlePlayer) GetNoOfPossibleSolutions() int {
	return player.possibleSolutions.Count()
}

func (player *WordlePlayer) GetPossibleSolutions() string {

	var wordsAsStrings []string
	for _, word := range player.possibleSolutions.Words {
		wordsAsStrings = append(wordsAsStrings, word.String)
	}

	return strings.Join(wordsAsStrings, ", ")
}
