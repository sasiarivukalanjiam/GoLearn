package main

import (
	"bufio"
	"fmt"
	"myapp/doctor"
	"os"
	"strings"
)

/*
func main() {
	fmt.Println("hello world")
}
*/

/*
func main() {
	sayHelloWorld("Hellow, world!!")
}*/

/*
func main() {
	var whatToSay string
	whatToSay = "Hello, World"
	sayHelloWorld(whatToSay)
}
*/

/*
func main() {
	//var whatToSay string
	//whatToSay = "Hello, World"
	whatToSay := "hello world"
	sayHelloWorld(whatToSay)
}
*/

/*
func main() {
	reader := bufio.NewReader(os.Stdin)

	//var whatToSay string
	whatToSay := doctor.Intro()
	fmt.Println(whatToSay)

	for {
		userInput, _ := reader.ReadString('\n')
		fmt.Println(userInput)
	}

}

*/

/*
func sayHelloWorld(whatToSay string) {
	fmt.Println(whatToSay)
}
*/

func main() {
	reader := bufio.NewReader(os.Stdin)

	//var whatToSay string
	whatToSay := doctor.Intro()
	fmt.Println(whatToSay)

	for {
		fmt.Print("-> ")
		userInput, _ := reader.ReadString('\n')

		userInput = strings.Replace(userInput, "\r\n", "", -1)
		response := doctor.Response(userInput)

		if userInput == "quit" {
			break
		} else {
			fmt.Println(response)
		}
	}

}
