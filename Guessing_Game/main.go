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
	newScanner := bufio.NewReader(os.Stdin)

	var number int64
	var isGuessedRight bool

	// seed to get different result every time
	rand.Seed(time.Now().UnixMilli())
	number = int64(rand.Intn(100) + 1)
	// fmt.Println("Number is", number)

	for i := 0; i < 10; i++ {
		fmt.Printf("You have got %d chance left\n", 10-i)
		fmt.Printf("Guess an integer between 1 and 100: ")

		input, err := newScanner.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		input = strings.TrimSpace(input)
		guessingNumber, err := strconv.ParseInt(input, 0, 64)

		if err != nil {
			log.Fatal(err)
		}

		if guessingNumber > number {
			fmt.Println("Guessing number is bigger")
		} else if guessingNumber < number {
			fmt.Println("Guessing number is lower")
		} else {
			isGuessedRight = true
			fmt.Println("Congratulations, you guessed it right!")
			break
		}
	}

	if !isGuessedRight {
		fmt.Println("Sorry, your chance is finished. Try again later!")
	}

}
