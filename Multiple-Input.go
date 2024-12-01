package main

import (
	"fmt"
)

func main() {

	var name string
	var is_Wizard bool

	fmt.Print("Enter your name & are you a Wizard: ")
	fmt.Scanf("%s %t", &name, &is_Wizard)
	fmt.Println(name, is_Wizard)

}
