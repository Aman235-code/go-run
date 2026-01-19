package main

import (
	"fmt"
	// "time"
)

// order struct
// type order struct {
// 	id string 
// 	amount float32
// 	status string
// 	createdAt time.Time
// }

// func newOrder(id string, amount float32, status string) *order {
// 	 myOrder := order{
// 		id: id,
// 		amount: amount,
// 		status: status,
// 	 } 

// 	 return &myOrder
// }

// // receiver type
// func (o *order) changeStatus(status string){
// 	o.status = status
// }

// func (o order) getAmount() float32{
// 	return o.amount
// }


func main(){

	language := struct {
		name string 
		isGood bool
	} {"golang", true}

	fmt.Println(language)

	// myOrder := newOrder("1", 30.60, "received")
	// fmt.Println(myOrder)
	// if you dont set any field value, default value is zero value
	// int = 0, float = 0, string = "", bool = false
	//  myOrder := order{
	// 	id: "1",
	// 	amount: 50.00,
	// 	status: "received",
	//  } 

	//  myOrder.changeStatus("confirmed")

	//  fmt.Println(myOrder)

	//  myOrder.createdAt = time.Now()

	//  fmt.Println(myOrder)
	// // fmt.Println(myOrder.status)

	// myOrder2 := order {
	// 	id: "2",
	// 	amount: 100,
	// 	status: "delivered",
	// 	createdAt: time.Now(),
	// }

	// fmt.Println(myOrder2)
}