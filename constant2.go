package main

import "fmt"

// The simple constant declaration ":=" below will throw an error because we are trying to reassign a value to a constant.
func main() {
	const name = "Ender"
	name = "Bean"
	fmt.Printf("%v: %T \n", name, name)
}
