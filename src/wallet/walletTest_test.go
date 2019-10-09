package test

import (
	"testing"
	"fmt"
)

func TestWallet(t *testing.T){
	wallet := Wallet{}
	fmt.Println("address of balance in Deposit is", &wallet.balance)
	wallet.Deposit(Bitcoin(10))
	got := wallet.Balance()
	want := 10

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}

}