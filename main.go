package main

import (
	"fmt"

	"github.com/seon0025/learngo/accounts"
)

func main() {
	account := accounts.NewAccount("seon")
	account.Deposit(10)
	fmt.Println(account)
}
