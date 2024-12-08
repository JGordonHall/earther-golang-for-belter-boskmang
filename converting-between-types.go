// Type Castering  = The name of the process of converting a value from one type to another is known as TypeCasting
// Data types can be converted to other data types, but this does not guarantee that the value will remain intact.

//package main

//import "fmt"

//func main() {
//	var i int = 90
//	var f float64 = float64(i)
//	fmt.Printf("%.2f\n", f)
//}

//package main

//import "fmt"

//func main() {
//	var f float64 = 45.89
//	var i int = int(f)
//	fmt.Printf("%v\n", i)
//}

// strconv package = The strconv package provides functions to convert a string to a number and vice-versa.
// Iota() converts integer to string, returns one value-string formed with the given integer.

package main

import (
	"fmt"
	"strconv"
)

func main() {
	var i int = 42
	var s string = strconv.Itoa(i) // conver int to string
	fmt.Printf("%q", s)

}
