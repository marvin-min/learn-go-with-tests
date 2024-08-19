package pointer_error

import (
	"errors"
	"fmt"
)

type Bitcoin int
type Wallet struct {
	balance Bitcoin
}

type Stringer interface {
	String() string
}

var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")
var ErrInvalidAmount = errors.New("cannot withdraw, invalid amount")

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}
func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

func (w *Wallet) Withdraw(amount Bitcoin) error {
	if amount < 0 {
		return ErrInvalidAmount
	} else if w.balance >= amount && amount >= 0 {
		w.balance -= amount
		return nil
	} else {
		return ErrInsufficientFunds
	}
}
