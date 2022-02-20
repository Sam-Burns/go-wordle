package words

type Wordlist struct {
	Words []Word
}

func (wl Wordlist) Count() int {
	return len(wl.Words)
}

func makeWordlist(words []Word) *Wordlist {
	return &Wordlist{words}
}
