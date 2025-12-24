package manager

import (
	"hotel-mgmt-lld/internal/domain"
	interfaces "hotel-mgmt-lld/internal/interface"
	"hotel-mgmt-lld/internal/usecase"
)

type PaymentMgr struct {
}

func NewPaymentMgr() *PaymentMgr {
	return &PaymentMgr{}
}

func (pm *PaymentMgr) GetPaymentStrategy(PaymentStrategy domain.PaymentStrategy) interfaces.IPayment {
	switch PaymentStrategy {
	case domain.CardPayment:
		return &usecase.CardPayment{}
	}
	return nil
}

func (pm *PaymentMgr) Pay(strategy interfaces.IPayment, amount int) {
	strategy.Pay(amount)
}
