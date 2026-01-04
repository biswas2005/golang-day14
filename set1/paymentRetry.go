package set1

import (
	"errors"
	"fmt"
)

type PaymentGateway interface {
	Pay(amount float64) error
}

var payErr = errors.New("Not sufficient amount")

type UPI struct {
	ID      string
	Balance float64
}

func (u *UPI) Pay(amount float64) error {

	if amount <= 0 || amount > u.Balance {
		return payErr
	}
	fmt.Printf("Amount %.2f paid by %s\n", amount, u.ID)
	u.Balance -= amount
	return nil
}

func PaymentRetry() {
	chance := 3
	upi := UPI{ID: "qwe123@xyz", Balance: 1000}

	var user float64

	for chance > 0 {
		fmt.Println("Enter amount to pay:")
		fmt.Scan(&user)
		if err := upi.Pay(user); err != nil {
			chance--
			fmt.Printf("Payment failed.You have %d Chance\n", chance)

			if chance == 0 {
				fmt.Println("Payment aborted.")
				break
			}
		} else {
			fmt.Println("Payment successful.")

			fmt.Printf("Available Balance:%.2f", upi.Balance)
			break

		}

	}

}
