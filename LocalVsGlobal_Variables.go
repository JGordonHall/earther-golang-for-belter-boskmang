package main

import (
	"fmt"
)

var Name string = "Gordon"

func main() {

	fmt.Println(Name)

}

//AI Refactored Example

package main

import "fmt"

// First main function logic
func firstMainLogic() {
    fmt.Println("First main logic")
}

// Second main function logic
func secondMainLogic() {
    fmt.Println("Second main logic")
}

func main() {
    firstMainLogic()
    secondMainLogic()
}