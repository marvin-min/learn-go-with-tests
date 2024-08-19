package pointer_error

import (
	"testing"
)

func TestWallet(t *testing.T) {
	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))

		assetBalance(wallet, Bitcoin(10), t)
	})

	t.Run("withdraw", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(20))
		wallet.Withdraw(Bitcoin(10))
		assetBalance(wallet, Bitcoin(10), t)
	})
	t.Run("withdraw insufficient funds", func(t *testing.T) {
		startBalance := Bitcoin(20)
		wallet := Wallet{startBalance}
		err := wallet.Withdraw(Bitcoin(100))
		assetError(t, err, ErrInsufficientFunds)
		assetBalance(wallet, startBalance, t)

	})

	t.Run("withdraw amount less than zero", func(t *testing.T) {
		startBalance := Bitcoin(20)
		wallet := Wallet{startBalance}
		err := wallet.Withdraw(Bitcoin(-10))
		assetError(t, err, ErrInvalidAmount)
		assetBalance(wallet, startBalance, t)
	})
}

func assetError(t *testing.T, got error, want error) {
	t.Helper()
	if got == nil {
		t.Error("wanted an error but didn't get one")
	}
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func assetBalance(wallet Wallet, want Bitcoin, t *testing.T) {
	t.Helper()
	got := wallet.Balance()
	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}
