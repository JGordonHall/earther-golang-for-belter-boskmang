/*Description: This is an example of a for loop in Golang. For Loops are used to iterate over a range of values, and can be used to execute a block of code multiple times.

package main

import "fmt"

func main() {
	for i := 1; i <= 5; i++ {
		fmt.Println(i * i)
	}
}
// Output: 1	4	9	16	25  This is the output of the code snippet
// Explanation: The for loop iterates from 1 to 5, and prints the square of each number.

// Description: This is a for Loop without initialization and post statement in Golang

package main

import "fmt"

func main() {
	i := 1
	for i <= 5 {
		fmt.Println(i * i)
		i += 1
	}
}
// output : 1	4	9	16	25  This is the output of the code snippet
// Explanation: The for loop iterates from 1 to 5, and prints the square of each number.

// Description: This is an example of a infite for loop in Golang. It is a loop that runs indefinitely until a break statement is encountered. It occurs when the loop condition is always true or when there is no condition specified.

package main

import "fmt"

func main() {
	sum := 0
	for {
		sum++ // repeated forever
	}
	fmt.Println(sum) // This line will never be reached
}

// Output: This line will never be reached
// Explanation: The for loop runs indefinitely, incrementing the sum variable until a break statement is encountered. However, since there is no break statement in this example, the program will run indefinitely and never reach the Println statement.*/
