package main

import "fmt"

type orderstatus string

const (
	Recived   orderstatus = "recived"
	confirmed             = "confirmed"
	prepared              = "prepared"
	delivered             = "delivered"
)

func changeOrderstatus(status orderstatus) {
	fmt.Println(status)
}

func main() {
	changeOrderstatus(Recived)
}
