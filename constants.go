package main

import "fmt"

func main() {
	const name = "Gordon"
	const is_Engineer = true
	const age = 34

	fmt.Printf("%v: %T \n", name, name)
	fmt.Printf("%v: %T \n", is_Engineer, is_Engineer)
	fmt.Printf("%v: %T", age, age)

}
