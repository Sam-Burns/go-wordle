# go-wordle

An application for solving Wordle puzzles. Currently, the project analyses wordlists taken from the New York Times'
Javascript, and presents the results. The actual player is still in progress.

## Playing the Game

To play the game, run `go run ./cmd/play-wordle.go SPARE`, or with any other five-letter word that is a valid solution
to a Wordle.

```
go run cmd/play-wordle.go SPARE
Wordle solution: SPARE

There are currently 2309 possible solutions
The next guess should be ARISE
Worst-case scenario for proposed guess is the feedback -----. Carry-over ratio for possible solutions list would be 7.23%
Guess number 1: ARISE
Feedback from guess was: YY-YG

There are currently 5 possible solutions [SCARE, SNARE, STARE, SHARE, SPARE]
The next guess should be CHANT
Worst-case scenario for proposed guess is the feedback Y-G--. Carry-over ratio for possible solutions list would be 20.00%
Guess number 2: CHANT
Feedback from guess was: --G--

There are currently 1 possible solutions [SPARE]
Guess number 3: SPARE
Feedback from guess was: GGGGG

Won the Wordle in 3 turns

```

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
well-known 5-letter words.)
- The 5 most commonly used letters in valid Wordle solutions are the letters of the word **ORATE**
- If you wish to continue with the letter frequency strategy, the best two words to play next are **SULCI**, and then
**HANDY**. You will then have played the 14 top letters in 3 turns.

```
go run ./cmd/analyse-wordlists.go

There are 12947 words that are valid guesses
There are 2309 words that are valid solutions

Frequencies of characters in solution wordlist:

E	10.65%
A	 8.45%
R	 7.77%
O	 6.52%
T	 6.31%
L	 6.20%
I	 5.80%
S	 5.79%
N	 4.96%
C	 4.11%
U	 4.04%
Y	 3.67%
D	 3.40%
H	 3.35%
P	 3.16%
M	 2.74%
G	 2.69%
B	 2.43%
F	 1.98%
K	 1.82%
W	 1.68%
V	 1.32%
Z	 0.35%
X	 0.32%
Q	 0.25%
J	 0.23%
```
