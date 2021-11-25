package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/eiannone/keyboard"
)

// https://github.com/eiannone/keyboard
// go get -u github.com/eiannone/keyboard

func main() {

	/*
		reader := bufio.NewReader(os.Stdin)
		for {
			fmt.Print("-> ")
			userInput, _ := reader.ReadString('\n')
			userInput = strings.Replace(userInput, "\r\n", "", -1)
			if userInput == "quit" {
				break
			} else {
				fmt.Println(userInput)
			}
		}
	*/

	if err := keyboard.Open(); err != nil {
		panic(err)
	}
	defer func() {
		_ = keyboard.Close()
	}()

	fmt.Println("Press ESC to quit")
	for i := 0; i < 10; i++ {
		char, key, err := keyboard.GetKey()
		if err != nil {
			panic(err)
		}
		fmt.Printf("You pressed: rune %q, key %X\r\n", char, key)
		if key == keyboard.KeyEsc {
			break
		}
	}
	fmt.Println("Exiting program")

	colour := make(map[int]string)
	colour[1] = "green"
	colour[2] = "red"
	colour[3] = "blue"
	colour[4] = "orange"
	colour[5] = "black"
	colour[6] = "pink"
	colour[7] = "red"
	colour[8] = "yellow"
	colour[9] = "violet"

	fmt.Println("choose color menu")
	fmt.Println("-----------------")
	fmt.Println("1 - green")
	fmt.Println("2 - red")
	fmt.Println("3 - blue")
	fmt.Println("4 - orange")
	fmt.Println("5 - black")
	fmt.Println("6 - pink")
	fmt.Println("7 - red")
	fmt.Println("8 - yellow")
	fmt.Println("9 - violet")
	fmt.Println("press Q to quit the program")

	for {
		char, _, err := keyboard.GetSingleKey()
		if err != nil {
			log.Fatal(err)
		}

		i, _ := strconv.Atoi(string(char))
		fmt.Println("value if i :", i)
		//t := fmt.Sprintf("actual choice %q", char)
		//t := fmt.Sprintf("actual choice %d", i)
		t := fmt.Sprintf("actual choice %s", colour[i])
		fmt.Println("you choose", t)
		//fmt.Println("you choose", char)

		if char == 'q' || char == 'Q' {
			break
		}
	}
	fmt.Println("exiting code")
}
