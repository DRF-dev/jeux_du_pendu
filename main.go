package main

import (
	"fmt"
	"os"
	"udemygo/jeux_du_pendu/dictionnaire"
	"udemygo/jeux_du_pendu/hangman"
)

func main() {

	err := dictionnaire.Load("words.txt")
	if err != nil {
		fmt.Printf("Erreur : %v\n", err)
		os.Exit(1)
	}

	g := hangman.NewGame(8, dictionnaire.PickWord())
	hangman.DrawWelcome()
	guess := ""
	for {
		hangman.Draw(g, guess)

		switch g.State {
		case "won", "lost":
			os.Exit(0)
		}

		l, err := hangman.ReadGuess()
		if err != nil {
			fmt.Printf("Erreur : %v\n", err)
			os.Exit(1) //Si os.Exit vaut 0 sa a marché si il vaut 1 sa a échoué
		}
		guess = l
		g.MakeAGuess(guess)
	}
}
