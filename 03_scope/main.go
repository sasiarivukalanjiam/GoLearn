package main

import (
	"myapp/packageone"
)

/*
var two = "package variable "

func main() {
	var one = "variable in main"

	fmt.Println(one)

	myFunc()

	newString := packageone.PublicVariable
	fmt.Println("from packageone", newString)
	packageone.Exported()
}

func myFunc() {
	var one = "one inside myFunc"
	fmt.Println(one)
	fmt.Println(two)
	//redeclare two variable will be variable shadowing
}


*/

var myVar = "package level variable"

func main() {
	var blockVar = "block level variable "
	packageone.PrintMe(myVar, blockVar)
}
