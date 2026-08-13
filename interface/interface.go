package main

import "fmt"

type paymenter interface {
	pay(amount float32)
}

type payment struct {
	gateway paymenter
}

func (p payment) makePayment(amount float32) {
	p.gateway.pay(amount)
}

type stripeGateway struct {
}

func (s stripeGateway) pay(amount float32) {
	fmt.Println("Payment done using stripe : ", amount)
}

type surjoPay struct {
}

func (s surjoPay) pay(amount float32) {
	fmt.Println("Payment done using surjoPay : ", amount)
}

func main() {
	var amount float32
	fmt.Println("Make your payment amount: ")
	fmt.Scanln(&amount)
	//stripe := stripeGateway{}
	surjo := surjoPay{}
	payment := payment{gateway: surjo}
	payment.makePayment(amount)

}
