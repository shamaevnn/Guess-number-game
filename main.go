package main

import (
	"errors"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func getRandomNumber(maxNumber int) int {
	// to always get random number we need to provide seed
	seed := time.Now().UnixNano()
	s1 := rand.NewSource(seed)
	r1 := rand.New(s1)

	randomNumber := r1.Intn(maxNumber)
	return randomNumber
}

func getParamsFromCommandLine() (map[string]int, error) {
	argsWithoutProg := os.Args[1:]

	if len(argsWithoutProg) != 2 {
		return nil, errors.New("only two arguments must be passed. The first one is maxNumber and second one is numberOfAttempts")
	}

	maxNumberString := argsWithoutProg[0]
	maxNumber, err := strconv.Atoi(maxNumberString)
	if err != nil {
		return nil, errors.New("maxNumber is not a number, try again")
	}

	numberOfAttemptsString := argsWithoutProg[1]
	numberOfAttempts, err := strconv.Atoi(numberOfAttemptsString)
	if err != nil {
		return nil, errors.New("numberOfAttempts is not a number, try again")
	}

	res := map[string]int{"maxNumber": maxNumber, "numberOfAttempts": numberOfAttempts}
	return res, nil
}

func main() {
	params, err := getParamsFromCommandLine()
	if err != nil {
		log.Fatal(err)
	}

	maxNumber := params["maxNumber"]
	numberOfAttempts := params["numberOfAttempts"]

	computerNumber := getRandomNumber(maxNumber)
	fmt.Printf("I thought of a number, try to guess it! You have %d attempts\n", numberOfAttempts)

	for numberOfAttempts > 0 {
		var input string
		_, err := fmt.Scanln(&input)
		if err != nil {
			println("Invalid input, try again")
			continue
		}

		inputNumber, err := strconv.Atoi(input)
		if err != nil {
			fmt.Printf("%s is not a number, try again\n", input)
			continue
		}

		if inputNumber > computerNumber {
			fmt.Printf("You number is greater, you have %d attempts\n", numberOfAttempts-1)
			numberOfAttempts -= 1
		} else if inputNumber < computerNumber {
			fmt.Printf("You number is less than computer number, you have %d attempts\n", numberOfAttempts-1)
			numberOfAttempts -= 1
		} else if inputNumber == computerNumber {
			println("Congratulations, you guessed it!")
			return
		}

		if numberOfAttempts == 0 {
			fmt.Printf("Game over. The computer thought of %d\n", computerNumber)
			return
		}
	}
}
