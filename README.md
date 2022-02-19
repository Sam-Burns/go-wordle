# go-wordle

An application for solving Wordle puzzles. Currently, the project analyses wordlists taken from the New York Times'
Javascript, and presents the results. The actual player is still being built.

## Wordlists

`./data/wordlist-valid-guesses.csv` and `./data/wordlist-valid-solutions.csv` are the input files. They are taken from
the New York Times website's Javascript.

You can very simply reverse engineer the Javascript by going to
[The NYT Wordle page](https://www.nytimes.com/games/wordle/index.html)
and checking what Javascript is being used. At time of writing, it is
https://www.nytimes.com/games/wordle/main.4d41d2be.js .

## Analysis of Wordlists

Run `go build analyse-wordlists.go` to run the analysis. You will need Go v1.17. The results are as follows:

- There are 12,497 words that _you_ are allowed to use as your _guess_
- Of these, there are 2,309 words that _they_ are allowed to use as the _solution_ to a Wordle. (These tend to be
well-known 5 letter words.)
- The 5 most commonly used letters in valid Wordle solutions are the letters of the word **AROSE**. The next 5 most 
common letters are the letters of the word **UNTIL**.
