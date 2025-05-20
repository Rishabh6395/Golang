package main

import "fmt"

// for -> only construct in go for looping

func main() {
	// while loop

	// i := 1
	// for i <= 3 {
	// 	fmt.Println(i)
	// 	i = i + 1
	// }

	// normal for loop

	// for i := 0; i < 3; i++ {
	// 	fmt.Println(i)
	// }

	// range

	for i := range 3 {
		fmt.Println(i)
	}

}
