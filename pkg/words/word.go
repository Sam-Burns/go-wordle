package words

import "errors"

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

	if len(wordString) != 5 {
		return nil, errors.New("total fuck-up")
	}

	return &Word{
		[]rune(wordString),
		wordString,
	}, nil
}
