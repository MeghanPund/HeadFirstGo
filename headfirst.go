// generate random num from 1-100 and store as target for player to guess
// prompt player to guess target #, store their response
// if guess < target, "Oops. Your guess was LOW"
// if > "Oops. Your guess was HIGH"
// otherwise, "Good job. You guessed it!"
// Allow 10 guesses
// Before each guess, tell how many guesses remaining
// if player fails all 10 guesses "Sorry. You didn't guess my number. It was: [target]."

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
	fmt.Println("I've chosen a random number between 1 and 100.\nCan you guess it?")
	fmt.Println(target)

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Your guess: ")
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	input = strings.TrimSpace(input)
	guess, err := strconv.Atoi(input)
	if err != nil {
		log.Fatal(err)
	}

	if guess == target {
		fmt.Println("Congrats! You guessed it.")
	} else if guess > target {
		fmt.Println("Your guess was too HIGH.")
	} else {
		fmt.Println("Your guess was too LOW.")
	}
}
