package main

import (
	"fmt"
	"time"
)

type order struct {
	id        string
	amount    float32
	status    string
	createdAt time.Time //nanosecond precision
}

// reciver type
func (o *order) chagestatus(status string) {
	o.status = status
}
func (o *order) getamount() float32 {
	return o.amount
}
func neworder(id string, amount float32) *order {
	myorder3 := order{
		id:     id,
		amount: amount,
		status: "pending",
	}
	return &myorder3
}

func main() {
	myorder := order{
		id:     "123",
		amount: 100,
		status: "pending",
		// createdAt: time.Now(),
	}
	myorder2 := order{
		id:     "343",
		amount: 434,
		status: "complete",
	}
	myorder3 := neworder("123", 100)
	fmt.Println(myorder3) 
	myorder.createdAt = time.Now() //.Add(time.Hour * 24 * 7)
	fmt.Println(myorder.createdAt)
	fmt.Println(myorder)
	fmt.Println(myorder2)
	fmt.Println(myorder.getamount())
	myorder.chagestatus("complete")
	fmt.Println(myorder)
}
