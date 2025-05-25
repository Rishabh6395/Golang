package main

import (
	"fmt"
	"maps"
)

// maps -> hash, object
func main() {
	// creating map

	// m := make(map[string]string)

	// setting an elements
	// m["name"] = "golang"
	// m["area"] = "backend"

	// get an element
	// IMP: if key does not exists in the map then it returns zero value

	// m := make(map[string]int)
	// m["age"] = 30
	// m["price"] = 50
	// fmt.Println(m["name"], m["area"])
	// fmt.Println(len(m))

	// delete(m, "price")
	// clear(m)
	// fmt.Println(m)

	// m := map[string]int{"price": 40, "phones": 3}

	// v, ok :=m["price"]
	// fmt.Println(v)

	// if ok {
	// 	fmt.Println("all ok")
	// }else{
	// 	fmt.Println("not ok")
	// }

	m := map[string]int{"price": 40, "phones": 3}
	m1 := map[string]int{"price": 40, "phones": 3}

	fmt.Println(maps.Equal(m, m1))

}
