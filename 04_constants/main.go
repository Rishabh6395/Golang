package main

import "fmt"

const age = 30

func main() {
	// const name = "GOlang"
	// const age = 30

	const (
		port = 5000
		host = "localhost"
	)

	fmt.Println(port, host)

	fmt.Println(age)
}
