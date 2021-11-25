package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

const prompt = "and press enter when ready"

func main() {
	/*
		// declare and assign
		var firstNumber int
		firstNumber = 2

		var secondNumber = 5

		subtraction := 7
	*/

	rand.Seed(time.Now().UnixNano())

	var firstNumber = rand.Intn(8) + 2  //2
	var secondNumber = rand.Intn(8) + 2 //5
	var subtraction = rand.Intn(8) + 2  //7
	var answer = firstNumber*secondNumber - subtraction
	playTheGame(firstNumber, secondNumber, subtraction, answer)

}

func playTheGame(firstNumber, secondNumber, subtraction, answer int) {
	reader := bufio.NewReader(os.Stdin)

	// welcome message
	fmt.Println("predict the number ")
	fmt.Println("-------------------")
	fmt.Println("")

	fmt.Println("enter a number between 1 and 10 ", prompt)

	reader.ReadString('\n')

	// play game
	fmt.Println("multiply your number by", firstNumber, prompt)
	reader.ReadString('\n')

	fmt.Println("multiply the number by ", secondNumber, prompt)
	reader.ReadString('\n')

	fmt.Println("divide by the number that you originally thought", prompt)
	reader.ReadString('\n')

	fmt.Println("now subtract ", subtraction, prompt)
	reader.ReadString('\n')

	//provide answer
	//answer = firstNumber*secondNumber - subtraction
	fmt.Println("the answer is ", answer)
}
