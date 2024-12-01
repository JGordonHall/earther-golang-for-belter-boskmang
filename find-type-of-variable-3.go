package main

import (
	"fmt"
	"reflect"
)

func main() {
	var grades int = 42
	var message string = "Black Mesa"

	fmt.Printf("variable grades=%v is of type %v \n", grades, reflect.TypeOf(grades))
	fmt.Printf("variable message='%v' is of type %v \n", message, reflect.TypeOf(message))

}
