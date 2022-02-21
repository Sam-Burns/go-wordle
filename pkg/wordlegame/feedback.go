package wordlegame

import "github.com/Sam-Burns/go-wordle/pkg/words"

type FeedbackColour int64

const (
	Green FeedbackColour = iota
	Yellow
	Grey
)

type Feedback struct {
	feedbackColours []FeedbackColour
}

func GetFeedback(solution *words.Word, guess *words.Word) *Feedback {

	feedbackColours := []FeedbackColour{Grey, Grey, Grey, Grey, Grey}

	for index, _ := range guess.Characters {
		feedbackColours[index] = getFeedbackColor(solution, guess, &index)
	}

	return &Feedback{feedbackColours}
}

func getFeedbackColor(solution *words.Word, guess *words.Word, index *int) FeedbackColour {

	if solution.ContainsCharacterAtIndex(guess.Characters[*index], *index) {
		return Green
	}
	if solution.ContainsCharacterAnywhere(guess.Characters[*index]) {
		return Yellow
	}

	return Grey
}

func (f Feedback) String() string {
	feedbackString := ""
	for _, colour := range f.feedbackColours {
		switch colour {
		case Grey:
			feedbackString += "-"
		case Yellow:
			feedbackString += "Y"
		case Green:
			feedbackString += "G"
		}
	}
	return feedbackString
}

func (f Feedback) Equals(another Feedback) bool {
	return f.String() == another.String()
}
