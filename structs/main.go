package main

import (
	"fmt"
	"time"
)

// order struct

type customer struct {
	name  string
	phone int
}

type order struct {
	id        string
	amount    float32
	status    string
	createdAt time.Time // nanosecond precision
	customer
}

// func newOrder(id string, amount float32, status string) *order {
// initial setup goes here..
// 	myOrder := order{
// 		id:     id,
// 		amount: amount,
// 		status: status,
// 	}
// 	return &myOrder
// }

// func (o *order) changeStatus(status string) {
// 	o.status = status
// }

// func (o order) getAmount() float32 {
// 	return o.amount
// }

func main() {
	newOrder := order{
		id:     "1",
		amount: 30,
		status: "pending",
		customer: customer{
			name:  "fahad",
			phone: 899,
		},
	}
	fmt.Println(newOrder)
	// language := struct {
	// 	name   string
	// 	isGood bool
	// }{"golang", true}

	// fmt.Println(language)
	// myOrder := newOrder("1", 44.3, "paid")
	// fmt.Println(myOrder)
	// myOrder := order{
	// 	id:     "1",
	// 	amount: 50.00,
	// 	status: "pending",
	// }
	// myOrder.changeStatus("paid")
	// myOrder.getAmount()

	// myOrder.createdAt = time.Now()
	// fmt.Println(myOrder)
}
