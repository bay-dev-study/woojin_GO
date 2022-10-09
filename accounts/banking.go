package accounts

import "fmt"

// Account struct
type Account struct {
	owner   string
	balance int
}

//대문자 퍼블릭 소문자 프라이빗

// NewAccount creates Account
func NewAccount(owner string) *Account {
	account := Account{owner: owner, balance: 0}
	return &account
}

//Deposit x amount on your account
func (a *Account) Deposit(amount int) {
	fmt.Println("gonna deposit", amount)
	a.balance += amount
}

func (a Account) Balance() int {
	return a.balance
}
