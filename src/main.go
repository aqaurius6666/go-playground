package main

import (
	"fmt"
	"log"

	guessgame "github.com/aquarius6666/go-playground/src/guess-game"
)

func main() {
	target := guessgame.GetTarget()
	success := false
	for i := 1; i < 10; i++ {
		guessStr, err := guessgame.GetGuess()
		if err != nil {
			log.Fatal(err)
		}

		guess, err := guessgame.ParseInt(guessStr)
		if err != nil {
			log.Fatal(err)
		}
		check := guessgame.Check(&target, &guess)
		if check == 0 {
			success = true
			break
		}
		if check > 0 {
			fmt.Println("You guessed too high!")
		}
		if check < 0 {
			fmt.Println("You guessed too low!")
		}
	}
	if !success {
		fmt.Printf("You lose! Secret number is %d\n", target)
	} else {
		fmt.Printf("You win! Secret number is %d\n", target)
	}
}
