/* Descripton: Logical operators in Golang "AND (&&)""
package main

import "fmt"

func main() {
	var x int = 10
	fmt.Println((x < 100) && (x < 200))
	fmt.Println((x < 300) && (x < 0))
}

Output:	true
		false

// Description: Logical operators in Golang "OR (||)""

package main

import "fmt"

func main() {
	var x int = 10
	fmt.Println((x < 0) || (x < 200))
	fmt.Println((x < 0) || (x > 200))
}

Output:	true
		false


Description: Logical operators in Golang "NOT (!)""

package main

import "fmt"

func main() {
	var x, y int = 10, 20
	fmt.Println(!(x > y))
	fmt.Println(!(true))
	fmt.Println(!(false))

}

Output:	true
		false
		true  */