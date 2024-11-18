package main

import (
	"errors"
	"fmt"
)

type Account struct {
	balance float64
	owner string
}

func NewAccount(owner string) (*Account) {
	return &Account{
		balance: 0,
		owner:   owner,
	}
}

func (a *Account) SetBalance(newBalance float64) error {
	if newBalance < 0 {
		return fmt.Errorf("balance can not be less than zero")
	}
	a.balance = newBalance
	return nil
}

func (a *Account) GetBalance() float64 {
	return a.balance
}

func (a *Account) Deposit(amount float64) error {
	if amount < 0 {
		return errors.New("you can not deposit negative money")
	}
	a.balance = a.balance + amount
	return nil
}

func (a *Account) Withdraw(amount float64) error {
	if amount > a.balance {
		return fmt.Errorf("insufficient funds: your balance %.2f", a.balance)
	}
	if amount < 0 {
		return fmt.Errorf("can not withdraw negative balance %.2f", a.balance)
	}
	a.balance = a.balance - amount
	return nil
}

func main(){
	account1 := NewAccount("Danial")
	account1.GetBalance()
	account1.Deposit(1000.0)
	
}