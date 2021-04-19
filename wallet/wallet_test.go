package main

import "testing"

func TestWallet(t *testing.T) {
	assertBalance := func(t testing.TB, wallet Wallet, want Bitcoin) {
		t.Helper()
		got := wallet.Balance()

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	}

	assertError := func(t testing.TB, got error, want string) {
		t.Helper()

		if got == nil {
			t.Fatal("wanted error but didn't get one")
		}

		if got.Error() != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	assertNotError := func(t testing.TB, err error) {
		if err != nil {
			t.Errorf("unexpected error: %q", err)
		}
	}

	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))

		assertBalance(t, wallet, Bitcoin(10))
	})

	t.Run("Withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}
		err := wallet.Withdraw(Bitcoin(10))

		assertNotError(t, err)
		assertBalance(t, wallet, Bitcoin(10))
	})

	t.Run("Insufficient founds", func(t *testing.T) {
		startingBalance := Bitcoin(10)
		wallet := Wallet{balance: startingBalance}
		err := wallet.Withdraw(Bitcoin(20))

		assertBalance(t, wallet, startingBalance)
		assertError(t, err, "insufficient founds")
	})
}
