package main

import (
	"fmt"
	"math/rand"
)

func Round(attemp int) bool {
	currentNumber := rand.Intn(100)
	//fmt.Println(currentNumber)
	var number int
	for i := 0; i < attemp; i++ {
		fmt.Print("Enter your guess: ")
		fmt.Scan(&number)
		if number == currentNumber {
			fmt.Printf("\nCongratulations! You guessed the correct number in 4 attempts.\n")
			return true
		}

		if number > currentNumber {
			fmt.Printf("\nIncorrect! The number is less than %d.\n", number)
			continue
		}

		fmt.Printf("\nIncorrect! The number is greater than %d.\n", number)
	}
	return false
}

func main() {
	fmt.Println("Welcome to the Number Guessing Game!\nI'm thinking of a number between 1 and 100.\nYou have 5 chances to guess the correct number.")
	fmt.Println("\nPlease select the difficulty level:\n1. Easy (10 chances)\n2. Medium (5 chances)\n3. Hard (3 chances)")
	fmt.Print("\nEnter your choice: ")
	var number int
	fmt.Scan(&number)
	fmt.Println()

	var attemp int
	switch number {
	case 1:
		fmt.Println("Great! You have selected the Easy difficulty level.")
		attemp = 10
	case 2:
		fmt.Println("Great! You have selected the Medium difficulty level.")
		attemp = 5
	case 3:
		fmt.Println("Great! You have selected the Hard difficulty level.")
		attemp = 3
	}
	fmt.Println("Let's start the game!")

	for i := 0; i < 5; i++ {
		fmt.Println(Round(attemp))
	}
}
