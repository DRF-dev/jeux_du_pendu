package hangman

import (
	"strings"
)

//Game => game
type Game struct {
	State        string
	Letters      []string //Les lettres à trouver = le mot
	FoundLetters []string //Les lettres trouvé
	UsedLetters  []string //Les lettres utilisées
	TurnsLeft    int      //Tour restant
}

//NewGame => newgame
func NewGame(turns int, word string) *Game {
	letters := strings.Split(strings.ToUpper(word), "") //On sépare word en un tableau a chaque fois qu'il y'a ""
	found := make([]string, len(letters))
	for i := 0; i < len(letters); i++ {
		found[i] = "_"
	}
	g := &Game{ //Nouvelle instance de Game
		State:        "",
		Letters:      letters,
		FoundLetters: found,
		UsedLetters:  []string{},
		TurnsLeft:    turns,
	}
	return g
}

//MakeAGuess => MakeAGuess
func (g *Game) MakeAGuess(guess string) {
	guess = strings.ToUpper(guess)

	if LetterInWord(guess, g.UsedLetters) {
		g.State = "alreadyGuess"
	} else if LetterInWord(guess, g.Letters) {
		g.State = "goodGuess"
		g.revealLetter(guess)
		if hasWon(g.Letters, g.FoundLetters) {
			g.State = "won"
		}
	} else if !LetterInWord(guess, g.Letters) {
		g.State = "badGuess"
		g.TurnsLeft--
		g.UsedLetters = append(g.UsedLetters, guess)
		if g.TurnsLeft <= 0 {
			g.State = "lost"
		}
	}
}

//LetterInWord => LetterInWord
func LetterInWord(guess string, letters []string) bool {
	for _, l := range letters {
		if l == guess {
			return true
		}
	}
	return false
}

func (g *Game) revealLetter(guess string) {
	g.UsedLetters = append(g.UsedLetters, guess)
	for i, l := range g.Letters {
		if l == guess {
			g.FoundLetters[i] = guess
		}
	}
}

func hasWon(letters []string, foundLetters []string) bool {
	for i := range letters {
		if letters[i] != foundLetters[i] {
			return false
		}
	}
	return true
}
