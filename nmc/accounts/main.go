package main

import (
	"fmt"
	"log"

	"github.com/minpeter/learngo/nmc/accounts"
)

func main() {
	account := accounts.NewAccount("minpeter")
	account.Deposit(100)

	fmt.Println(account.String())

	err := account.Withdraw(30)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(account.String())

	account.ChangeOwner("minpeter2")

	fmt.Println(account.String())

}
