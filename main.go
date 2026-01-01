package main

import (
	"fmt"
	"math/rand"
)

// PLANNING...

// Welcome to the Number Guessing Game!
// I'm thinking of a number between 1 and 100.
// You have 5 chances to guess the correct number.

// Please select the difficulty level:
// 1. Easy (10 chances)
// 2. Medium (5 chances)
// 3. Hard (3 chances)

// Enter your choice: 2

// Great! You have selected the Medium difficulty level.
// Let's start the game!

// Enter your guess: 50
// Incorrect! The number is less than 50.

// Enter your guess: 25
// Incorrect! The number is greater than 25.

// Enter your guess: 35
// Incorrect! The number is less than 35.

// Enter your guess: 30
// Congratulations! You guessed the correct number in 4 attempts.

func main() {
	fmt.Println("Welcome to the Number Guessing Game!\nI'm thinking of a number between 1 and 100.\nYou have 5 chances to guess the correct number.")
	randomnumber := rand.Intn(100) + 1

	fmt.Println("rans: ", randomnumber)
	var usernumber int
	var lavel int
	var maxChances int

	fmt.Println("Please select the difficulty level \n1. Easy (10 chances)\n2. Medium (5 chances)\n3. Hard (3 chances)")

	fmt.Scan(&lavel)

	switch lavel {
		case 1:
			maxChances = 10
		case 2:
			maxChances = 5
		default:
			maxChances = 3
	}
	
	for round := 1; round <= maxChances; round++ {
		fmt.Printf("\nAttempt %d/%d - Enter the guess: ", round, maxChances)
		fmt.Scan(&usernumber)

		if usernumber == randomnumber {
			fmt.Println("🎉 Congratulations! You guessed the correct number!")
			return
		}

		if usernumber > randomnumber {
			fmt.Println("Incorrect! The number is less than", usernumber)
		} else {
			fmt.Println("Incorrect! The number is greater than", usernumber)
		}
	}
	fmt.Println("\n😢 You've used all chances.")
	fmt.Println("The correct number was:", randomnumber)
}


