package accounts

import (
	"fmt"

	"github.com/wes-santos/alura-golang-study/bank/customers"
)

type CheckingAccount struct {
	Holder        customers.Holder
	AgencyNumber  int
	AccountNumber int
	balance       float64
}

func (c *CheckingAccount) Withdraw(amount float64) {
	canWithdraw := amount <= c.balance && amount > 0
	if canWithdraw {
		c.balance -= amount
		return
	}
	fmt.Printf(
		"You current balance is %v. You can't withdraw %v.",
		c.balance,
		amount,
	)
}

func (c *CheckingAccount) Deposite(amount float64) {
	if amount < 0 {
		fmt.Println("Deposite failed. The amount should be greater than zero.")
		return
	}

	c.balance += amount
}

func (c *CheckingAccount) Transfer(
	amount float64,
	destination *CheckingAccount,
) {
	if amount < 0 {
		fmt.Println("The transfer couldn't be done. Amount informed is less than zero.")
		return
	}

	if amount > c.balance {
		fmt.Println("The transfer couldn't be done. Amount informed is larger than account balance.")
		return
	}

	c.Withdraw(amount)
	destination.Deposite(amount)
}

func (c *CheckingAccount) GetBalance() float64 {
	return c.balance
}
