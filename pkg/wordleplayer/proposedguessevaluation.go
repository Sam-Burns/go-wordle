package wordleplayer

import (
	"fmt"
	"github.com/Sam-Burns/go-wordle/pkg/wordlegame"
	"github.com/Sam-Burns/go-wordle/pkg/words"
)

type ProposedGuessEvaluation struct {
	SizeOfCurrentShortlist           int
	ProposedGuess                    words.Word
	potentialFeedbackCounts          map[string]int
	worstCaseScenarioFeedbackString  string
	worstCaseShortlistCarryOverRatio float64
}

func MakeProposedGuessEvaluation(proposedGuess words.Word, sizeOfCurrentShortlist int) ProposedGuessEvaluation {
	return ProposedGuessEvaluation{sizeOfCurrentShortlist, proposedGuess, make(map[string]int), "", 0.0}
}

func (proposedGuessEvaluation ProposedGuessEvaluation) AddPossibleOutcome(possibleSolution words.Word, feedback wordlegame.Feedback) {
	proposedGuessEvaluation.potentialFeedbackCounts[feedback.String()] += 1
}

func (proposedGuessEvaluation *ProposedGuessEvaluation) GetWorstCaseScenarioFeedbackString() string {
	if proposedGuessEvaluation.worstCaseScenarioFeedbackString == "" {
		proposedGuessEvaluation.calculate()
	}
	return proposedGuessEvaluation.worstCaseScenarioFeedbackString
}

func (proposedGuessEvaluation *ProposedGuessEvaluation) getWorstCaseShortlistCarryOverRatio() float64 {
	if proposedGuessEvaluation.worstCaseShortlistCarryOverRatio == 0.0 {
		proposedGuessEvaluation.calculate()
	}
	return proposedGuessEvaluation.worstCaseShortlistCarryOverRatio
}

func (proposedGuessEvaluation *ProposedGuessEvaluation) calculate() {
	worstCaseScenario := struct {
		feedbackString string
		count          int
	}{
		"",
		0,
	}

	for potentialFeedbackString, potentialFeedbackCount := range proposedGuessEvaluation.potentialFeedbackCounts {

		if potentialFeedbackCount > worstCaseScenario.count {
			worstCaseScenario = struct {
				feedbackString string
				count          int
			}{
				potentialFeedbackString,
				potentialFeedbackCount,
			}
		}

	}

	proposedGuessEvaluation.worstCaseShortlistCarryOverRatio = float64(worstCaseScenario.count) / float64(proposedGuessEvaluation.SizeOfCurrentShortlist)
	proposedGuessEvaluation.worstCaseScenarioFeedbackString = worstCaseScenario.feedbackString
}

func (proposedGuessEvaluation ProposedGuessEvaluation) GetWorstCaseShortlistCarryOverRatioString() string {
	return fmt.Sprintf("%.2f", 100*proposedGuessEvaluation.getWorstCaseShortlistCarryOverRatio()) + "%"
}

func (proposedGuessEvaluation ProposedGuessEvaluation) IsNullEvaluation() bool {
	return proposedGuessEvaluation.worstCaseShortlistCarryOverRatio == 1.0
}
