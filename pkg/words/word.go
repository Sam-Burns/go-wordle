package words

import (
	"errors"
	"regexp"
)

type Word struct {
	Characters []rune
	String     string
}

func (w Word) ToString() (stringRepresentation string) {
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

func (w Word) ContainsCharacterAtIndex(needle rune, index int) bool {
	return w.Characters[index] == needle
}

func (w Word) ContainsCharacterAnywhere(needle rune) bool {
	for _, haystackCharacter := range w.Characters {
		if needle == haystackCharacter {
			return true
		}
	}
	return false
}

func (w Word) Equals(another *Word) bool {
	for index, char := range w.Characters {
		if another.Characters[index] != char {
			return false
		}
	}
	return true
}
