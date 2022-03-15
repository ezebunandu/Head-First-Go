// Simple number guessing game
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

	success := false
	for guesses := 0; guesses < 10; guesses++ {
		fmt.Print("Guess a number: ")
		reader := bufio.NewReader(os.Stdin)
		input, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		input = strings.TrimSpace(input)
		guessed_number, err := strconv.Atoi(input)
		if err != nil {
			log.Fatal(err)
		}
		if guessed_number > target {
			fmt.Println("You guessed", guessed_number, ": which is too high!")
		} else if guessed_number < target {
			fmt.Println("You guessed", guessed_number, ": that is too low")
		} else {
			success = true
			fmt.Println("You guessed right!")
			break
		}
		guesses_left := 9 - guesses
		fmt.Println("You have", guesses_left, "guesses left")
		fmt.Println("======================================")
	}
	if !success {
		fmt.Println("Sorry you didn't get the guess. It was: ", target)
	}

}
