package usecase

import "fmt"

type CardPayment struct {
}

func (cp *CardPayment) Pay(amount int) {
	fmt.Printf("Payment for amount : %v is done via card\n", amount)
}
