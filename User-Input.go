// package main

//import "fmt"

//func main() {

//	var name string
//	fmt.Print("Enter your name:")
//	fmt.Scanf("%s", &name)
//	fmt.Println("Hey there,", name)
//}

//package main

//import (
//	"fmt"
//)

//func main() {
//
//	var name string
//	var is_Engineer bool
//
//	fmt.Print("Enter you name & you are an Engineer:")
//	fmt.Scanf("%s %t", &name, &is_Engineer)
//	fmt.Println(name, is_Engineer)
//}

// Multiple Input

package main

import (
	"fmt"
)

func main() {
	var a string
	var b int

	fmt.Print("Enter a string and a number: ")

	count, err := fmt.Scanf("%s %d", &a, &b)

	fmt.Println("count :", count)
	fmt.Println("error: ", err)
	fmt.Println("a: ", a)
	fmt.Println("b: ", b)
}
