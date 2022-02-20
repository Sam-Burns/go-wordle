package words

import (
	"errors"
	"regexp"
)

type Word struct {
	Characters []rune
	String     string
}

func (w Word) toString() (stringRepresentation string) {
	for _, v := range w.Characters {
		stringRepresentation += string(v)
	}
	return
}

func MakeWord(wordString string) (*Word, error) {

	matchesValidationRegex, _ := regexp.MatchString("^[A-Z]{5}$", wordString)

	if !matchesValidationRegex {
		return nil, errors.New("String '" + wordString + "' not a valid word")
	}

	return &Word{
		[]rune(wordString),
		wordString,
	}, nil
}
