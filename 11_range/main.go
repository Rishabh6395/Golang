package main

import "fmt"

func main() {
	// nums := []int{6, 7, 8}

	// sum := 0
	// for i, num := range nums {
	// 	fmt.Println(num, i)
	// 	// sum = sum + num

	// }
	// fmt.Println(sum)

	// m := map[string]string{"fname": "john", "lname": "doe"}

	// for k, v:= range m{
	// 	fmt.Println(k, v)
	// }

	// unicode code point rune
	// starting byte of rune
	// 255 -> 1 byte, 2 byte
	for i, c := range "golang" {
		fmt.Println("in index", i, "with char", string(c), "unicode:", c)
	}

}
