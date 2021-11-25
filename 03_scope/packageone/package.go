package packageone

import "fmt"

/*
var privateVariable = "its private"
var PublicVariable = "its public"

func notExported() {

}

func Exported() {

}


*/

var PackageVar = "package level variable "

func PrintMe(s1, s2 string) {
	fmt.Println(s1, s2, PackageVar)
}
