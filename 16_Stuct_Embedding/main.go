package main

import (
	"fmt"
	"time"
)

type customer struct {
	name  string
	phone string
}

type order struct {
	id        string
	amount    float32
	status    string
	createdAt time.Time // nanosecond precision
	customer
}

func main() {

	newCustomer := customer{
		name:  "jhon",
		phone: "1234567",
	}
	newOrder := order{
		id:       "1",
		amount:   30,
		status:   "recived",
		customer: newCustomer,
	}
	newOrder.customer.name = "robin"
	fmt.Println(newOrder)
}
