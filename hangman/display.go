package hangman

import "fmt"

//DrawWelcome => Image bienvenue, (tapper sur google patorjk software taag)
func DrawWelcome() {
	fmt.Println(`
			____.                         .___      __________                   .___     
			|    | ____  __ _____  ___   __| _/_ __  \______   \ ____   ____    __| _/_ __ 
			|    |/ __ \|  |  \  \/  /  / __ |  |  \  |     ___// __ \ /    \  / __ |  |  \
		/\__|    \  ___/|  |  />    <  / /_/ |  |  /  |    |   \  ___/|   |  \/ /_/ |  |  /
		\________|\___  >____//__/\_ \ \____ |____/   |____|    \___  >___|  /\____ |____/ 
					\/            \/      \/                      \/     \/      \/      
	`)
}

//Draw => Dessin état de la partie
func Draw(g *Game, guess string) {
	drawTurns(g.TurnsLeft)
	drawState(g, guess)
}

//drawTurns
func drawTurns(l int) {
	var draw string
	switch l {
	case 0:
		draw = `Looser!!`
	case 1:
		draw = `Looser!`
	case 2:
		draw = `Looser`
	case 3:
		draw = `Loose`
	case 4:
		draw = `Loos`
	case 5:
		draw = `Loo`
	case 6:
		draw = `Lo`
	case 7:
		draw = `L`
	case 8:
		draw = ``

	}
	fmt.Println(draw)
}

//drawState
func drawState(g *Game, guess string) {
	fmt.Println("Guessed :")
	drawLetters(g.FoundLetters)

	fmt.Println("\nUsed :")
	drawLetters(g.UsedLetters)

	switch g.State {
	case "goodGuess":
		fmt.Println("\nGood guess")
	case "alreadyGuess":
		fmt.Printf("Letter %v was already used", guess)
	case "badGuess":
		fmt.Printf("Bad guess %v\n", guess)
	case "lost":
		fmt.Println("You lost :(")
	case "won":
		fmt.Println("You won :)")
	}
}

func drawLetters(l []string) {
	for _, c := range l {
		fmt.Printf("%v ", c)
	}
}
