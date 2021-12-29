package accounts

import (
	"errors"
	"fmt"
)

// Account struct
type Account struct {
	owner string
	balance int
}

// NewAccount creates a new account
func NewAccount(owner string) *Account {
	account := Account{owner: owner, balance: 0}
	return &account
}

//Deposit returns the description of the account
func (a *Account) Deposit(amount int) {
	a.balance += amount
}

//Balance returns the balance of the account
func (a Account) Balance() int {
	return a.balance
}

var errNoMoney = errors.New("Can't withdraw, not enough money")

//Withdraw returns the description of the account
func (a *Account) Withdraw(amount int) error {
	if a.balance < amount {
		return errNoMoney
	}
	a.balance -= amount
	return nil
}

//ChangeOwner changes the owner of the account
func (a *Account) ChangeOwner(newOwner string) {
	a.owner = newOwner
}

func (a Account) Owner() string {
	return a.owner
}

// like python class in __str__()

// String returns the description of the account
func (a Account) String() string {
	return fmt.Sprint(a.Owner(), ":", a.Balance())
}