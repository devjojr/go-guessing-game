package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	var userName string

	fmt.Println("Welcome to the number guessing game.")
	fmt.Print("What is your name: ")
	fmt.Scanln(&userName)

	correctNumber := rand.Intn(100) + 1
	numGuesses := 5
	guess := 0

	fmt.Printf("Hello %s, the rules are simple. You have to guess a number between 1 and 100. The number of guesses you have to select the correct number is %d.\n", userName, numGuesses)

	var guessMade int

	for guess < numGuesses {
		fmt.Print("What is your guess? ")
		fmt.Scanln(&guessMade)

		switch {
		case guessMade < correctNumber:
			fmt.Println("That guess is too low, try again.")
		case guessMade > correctNumber:
			fmt.Println("That guess is too high, try again.")
		default:
			fmt.Printf("Congrats %s!!! You guessed the correct number: %d. Your guess power level is over 9000!!!\n", userName, correctNumber)
			return
		}
		guess++
	}
	fmt.Printf("Oh no %s, you didn't guess the correct number :( the correct number was %d.\n", userName, correctNumber)
}
