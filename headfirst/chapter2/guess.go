// guess challenges palyer to guess random number.
package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	seconds := time.Now().Unix()
	rand.Seed(seconds)
	target := rand.Intn(100) + 1
	fmt.Println("I have choosen a random number between 1 and 100.")
	fmt.Println("Can you guess it?")
	

	reader := bufio.NewReader(os.Stdin)

	success := false

	for guesses:=0; guesses<10; guesses++ {
		fmt.Println("You have", 10-guesses, "guesses left.")


		fmt.Print("Make a guess: ")
		input, err := reader.ReadString('\n')

		if err != nil {
			log.Fatal(err)
		}

		input = strings.TrimSpace(input)
		guess, err := strconv.Atoi(input)

		if err != nil {
			log.Fatal(err)
		}

		if guess < target {
			fmt.Println("Guess was low!")
		} else if guess > target {
			fmt.Println("Guess was high!")
		} else {
			success=true
			fmt.Println("Good job you guessed it!")
			break
		}

	}

	if !success {
		fmt.Println("Sorry, You didn't guess my number. It was:", target)
	}

	
}