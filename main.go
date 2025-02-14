package main

import (
	"fmt"
	"math/rand"
	"time"
)

func selectDifficulty() int {
	fmt.Println("Please select the difficulty level:")
	fmt.Println("1. Easy 🙂 (10 chances)")
	fmt.Println("2. Medium 😊 (5 chances)")
	fmt.Println("3. Hard 🫡 (3 chances)")
	fmt.Println()
	
	var difficulty int
	fmt.Print("Select 1 / 2 / 3 : ")
	fmt.Scanln(&difficulty)
	fmt.Println()

	if difficulty < 1 || difficulty > 3 {
		difficulty = 1
		fmt.Println("🫤 Invalid selection! Defaulting to Easy mode.")
		fmt.Println()
	}

	return difficulty
}

func getChances(x int) int {
	difficulty := x
	var chances int

	switch difficulty {
	case 1:
		chances = 10
		fmt.Println("Great! Your difficulty level is Easy (10 chances).")
	case 2:
		chances = 5
		fmt.Println("Great! Your difficulty level is Medium (5 chances).")
	case 3:
		chances = 3
		fmt.Println("Great! Your difficulty level is Hard (3 chances).")
	}

	fmt.Println("🚩 Let's start the game!")
	fmt.Println()

	return chances
}


func check(x int, y int){
	chances := x
	random := y
	cnt := 0

	for (chances > 0) {
		var guess int
		fmt.Print("🤔 Enter your guess : ")
		fmt.Scanln(&guess)

		cnt++
		
		if guess >= 1 && guess <= 100 {
			if guess > random {
				fmt.Println("😬 Too High!")
			} else if guess < random {
				fmt.Println("☹️ Too Low!")
			} else {
				fmt.Printf("🎉 Congrats! You guessed the correct number in %d attemps!", cnt)
				fmt.Println()
				return
			}
		} else {
			fmt.Println("🫤 Invaild guess! Guess between 1 and 100!")
		}
		
		chances--
		
		if chances == 0 {
			fmt.Printf("😓 You are ran out of chances! The correct number was %d.", random)
			fmt.Println()
		}
	}
}

func generateRandomNumber() int {
	seed := time.Now().UnixNano()       // Get current time in nanoseconds as seed
	source := rand.NewSource(seed)      // Create a new random source
	randomGen := rand.New(source)       // Create a new random number generator
	random := randomGen.Intn(100) + 1   // Generates a random number between 1 and 100
	return random
}

func init() {
	fmt.Println("Welcome to the Number Guessing Game!")
	fmt.Println("I'm thinking of a number between 1 and 100.")
	fmt.Println("Your task is to guess the correct number.")
	fmt.Println()
}

func main() {
	difficulty := selectDifficulty()
	random := generateRandomNumber()
	chances := getChances(difficulty)
	check(chances, random)
}