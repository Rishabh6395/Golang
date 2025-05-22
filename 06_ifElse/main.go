package main

import "fmt"

func main() {
	// age := 16

	// if age >= 18 {
	// 	fmt.Println("Person is an adult")
	// } else {
	// 	fmt.Println("Person is not an adult")
	// }

	// if age >= 18 {
	// 	fmt.Println("Person is an adult")
	// } else if age >= 12 {
	// 	fmt.Println("Person is teen ager")
	// } else {
	// 	fmt.Println("Person is kid")
	// }

	// var role = "admin"
	// var hasPermissions = false

	// if role == "admin" && hasPermissions {
	// 	fmt.Println("Yes")
	// }

	// we can decalre a variable inside if construct
	if age := 15; age >= 18 {
		fmt.Println("Person is an adult", age)
	} else if age >= 12 {
		fmt.Println("Person is teenager", age)
	}

	// go does not have ternary, you will have to use normal if else

}
