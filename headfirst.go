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
	// grabs current time, which always changes; Unix converts to int
	seconds := time.Now().Unix()
	// uses the time to seed random num generator
	rand.Seed(seconds)
	// returns int between 0 and chosen val -1; adding +1 makes it btwn 1 and chosen val
	target := rand.Intn(100) + 1
	fmt.Println("I've chosen a random number between 1 and 100.\nCan you guess it?")
	fmt.Println(target)

	// reads keyboard input
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Your guess: ")
	// captures user input up to when they press Enter
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	// removes newline char
	input = strings.TrimSpace(input)
	// convwert input str to int
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
