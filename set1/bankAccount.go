package set1

import (
	"errors"
	"fmt"
)

// Account interface with multiple methods
type Account interface {
	Deposit(amount float64) error
	Withdraw(amount float64) error
}

type Statement interface {
	Display() float64
}

var (
	depoErr = errors.New("Amount must be greater than zero")
	withErr = errors.New("Insufficient Balance.")
)

type Bank struct {
	HolderName string
	Balance    float64
}

func (d *Bank) Deposit(amount float64) error {
	if amount <= 0 {
		return depoErr
	} else {
		d.Balance += amount
		fmt.Printf("Amount %.2f Deposited to %s\nUpdated Balance:%.2f", amount, d.HolderName, d.Balance)
	}
	return nil
}

func (w *Bank) Withdraw(amount float64) error {
	if amount > w.Balance {
		return withErr
	} else {
		w.Balance -= amount
		fmt.Printf("\nAmount %.2f withdrawn from %s\nUpdated Balance:%.2f", amount, w.HolderName, w.Balance)
	}
	return nil
}

func (s *Bank) Display() float64 {
	fmt.Printf("\nBalance:%.2f", s.Balance)
	return 0
}

type Bank1 struct {
	HolderName string
	Balance    float64
}

func (d *Bank1) Deposit(amount float64) error {
	if amount <= 0 {
		return depoErr
	} else {
		d.Balance += amount
		fmt.Printf("\nAmount %.2f Deposited to %s\nUpdated Balance:%.2f", amount, d.HolderName, d.Balance)
	}
	return nil
}

func (w *Bank1) Withdraw(amount float64) error {
	if amount > w.Balance {
		return withErr
	} else {
		w.Balance -= amount
		fmt.Printf("\nAmount %.2f withdrawn from %s\nUpdated Balance:%.2f", amount, w.HolderName, w.Balance)
	}
	return nil
}

func BankAccount() {
	b := Bank{HolderName: "abc", Balance: 1000}
	b1 := Bank1{HolderName: "xyz", Balance: 2000}

	if err := b.Deposit(200); err != nil {
		fmt.Println("Failed to Deposit.")
	}
	if err := b.Withdraw(400); err != nil {
		fmt.Println("Failed to Withdraw")
	}
	if err := b.Display(); err != 0 {
		fmt.Println("Unable to fetch Balance")
	}
	if err := b1.Deposit(30); err != nil {
		fmt.Println("Failed to deposit")
	}
	if err := b1.Withdraw(100); err != nil {
		fmt.Println("Failed to Withdraw")
	}
}
