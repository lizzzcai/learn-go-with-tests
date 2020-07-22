package wallet

import (
	"errors"
	"fmt"
)

type Bitcoin int

// Implement Stringer interface so that we can customize our string output of Bitcoin.
type Stringer interface {
	String() string
}

// This interface is defined in the fmt package and lets you define how your type is printed when used with the %s format string in prints.
func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

type Wallet struct {
	balance Bitcoin
}

func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

var ErrInsufficientFunds = errors.New("cannot widthdraw, insufficient funds")

func (w *Wallet) Withdraw(amount Bitcoin) error {
	if amount > w.balance {
		return ErrInsufficientFunds
	}
	w.balance -= amount
	return nil
}
