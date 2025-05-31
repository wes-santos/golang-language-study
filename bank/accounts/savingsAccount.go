package accounts

import (
	"fmt"

	"github.com/wes-santos/alura-golang-study/bank/customers"
)

type SavingsAccounts struct {
	Holder                                 customers.Holder
	AgencyNumber, AccountNumber, Operation int
	balance                                float64
}

func (c *SavingsAccounts) Withdraw(amount float64) {
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

func (c *SavingsAccounts) Deposite(amount float64) {
	if amount < 0 {
		fmt.Println("Deposite failed. The amount should be greater than zero.")
		return
	}

	c.balance += amount
}

func (c *SavingsAccounts) GetBalance() float64 {
	return c.balance
}
