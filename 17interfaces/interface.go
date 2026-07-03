package main

import "fmt"

type paymentgateway interface {
	pay(amount float32)
	refund(amount float32, account string)
}

type payment struct {
	gateway paymentgateway
}

func (p payment) makepayment(amount float32) {
	p.gateway.pay(amount)
}

type razorpay struct{}

func (r razorpay) pay(amount float32) {
	fmt.Println("makeing payment using razorpay", amount)
}

type stripe struct{}

func (s stripe) pay(amount float32) {
	fmt.Println("makeing payment using stripe", amount)
}

type paypal struct{}

func (p paypal) pay(amount float32) {
	fmt.Println("makeing payment using paypal", amount)
}

func (p paypal) refund(amount float32, account string) {
	fmt.Println("refund payment using paypal", amount, account)
}

func main() {
	paypalPaymentGw := paypal{}
	newpayment := payment{
		gateway: paypalPaymentGw,
	}
	newpayment.makepayment(100)
}
