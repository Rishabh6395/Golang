package main

import "fmt"

// func changeNum(num int) {
// 	num = 5
// 	fmt.Println("In changeNum", num)
// }

// by reference:
func changeNum(num *int){
	// defrenece 
	*num = 5
	fmt.Println("In changeNum", *num)
}

func main() {
	num := 1
	changeNum(&num)

	// fmt.Println("Memory address", &num)
	fmt.Println("After changeNum in main",num)
}