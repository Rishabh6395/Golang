package main // Declare the package name

import "fmt" // Import the fmt package for formatted I/O

// counter returns a closure function that increments and returns a counter
func counter() func() int {
	var count int = 0 // Initialize a local variable 'count' to 0

	return func() int { // Return an anonymous function (closure)
		count += 1   // Increment 'count' by 1
		return count // Return the updated value of 'count'
	}
}

func main() {
	increment := counter() // Call counter() and assign the returned function to 'increment'

	fmt.Println(increment()) // Call 'increment', prints 1
	fmt.Println(increment()) // Call 'increment' again, prints 2
	fmt.Println(increment()) // Call 'increment' again, prints 3
}
