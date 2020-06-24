package hangman

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

//ReadGuess => on lit la lettre écrite pas le user
func ReadGuess() (guess string, err error) {
	valid := false
	for !valid {
		fmt.Print("What is your letter ? ")
		guess, err = reader.ReadString('\n') //On lit notre reader dès que l'on va à la ligne = click sur entrée
		if err != nil {
			return guess, err
		}
		guess = strings.TrimSpace(guess) //Supprime les espaces
		if len(guess) != 1 {
			fmt.Printf("Element %v invalide car nombre de lettres %v invalide ! N'écrivez qu'une lettre\n", guess, len(guess))
			continue //On va direct en fin de boucle, en gros valid est encore égal à false et on repart au début de la boucle
		}
		valid = true
	}
	return guess, nil
}
