package words

type WordList struct {
	Words []Word
}

func (wl WordList) Count() int {
	return len(wl.Words)
}

func makeWordlist(words []Word) *WordList {
	return &WordList{words}
}
