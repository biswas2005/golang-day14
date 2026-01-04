package set1

import (
	"errors"
	"fmt"
)

type Paymenter interface {
	Pay(amount float64) error
}

var paymentErr = errors.New("Invalid amount.")

type CreditCard struct {
	CardNo string
}

func (c CreditCard) Pay(amount float64) error {
	if amount <= 0 {
		return paymentErr
	}
	fmt.Printf("Amount %.2f has been paid by %s\n", amount, c.CardNo)
	return nil
}

type PayPal struct {
	Number string
}

func (p PayPal) Pay(amount float64) error {
	if amount <= 0 {
		return paymentErr
	}
	fmt.Printf("Amount %.2f paid by %s\n", amount, p.Number)
	return nil
}

func Payments() {
	cc := CreditCard{CardNo: "1234-5678-9012"}
	pp := PayPal{Number: "0987654321"}

	if err := cc.Pay(100); err != nil {
		fmt.Println("Payment failed")
	}
	if err := pp.Pay(0); err != nil {
		fmt.Println("Payment failed")
	}

}
