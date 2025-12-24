package usecase

import "fmt"

type UpiPayment struct {
}

func (cp *UpiPayment) Pay(amount int) {
	fmt.Printf("Payment for amount : %v is done via upi\n", amount)
}
