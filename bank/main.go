package main

import (
	"fmt"

	"github.com/wes-santos/alura-golang-study/bank/accounts"
	"github.com/wes-santos/alura-golang-study/bank/customers"
)

func payBill(account AccountModel, billValue float64) {
	account.Withdraw(billValue)
}

type AccountModel interface {
	Withdraw(amount float64)
}

func main() {
	firstCheckingAccount := accounts.CheckingAccount{
		Holder: customers.Holder{
			Name:       "Weslley",
			CPF:        "424.424.424-42",
			Ocuppation: "Quantitative Developer",
		},
		AgencyNumber:  042,
		AccountNumber: 424242,
	}
	firstCheckingAccount.Deposite(208_000)

	secondCheckingAccount := accounts.CheckingAccount{
		Holder: customers.Holder{
			Name:       "Michael",
			CPF:        "424.424.424-43",
			Ocuppation: "Be a cat.",
		},
		AgencyNumber:  042,
		AccountNumber: 424243,
	}

	fmt.Println("Balance before withdraw: ", firstCheckingAccount.GetBalance())
	firstCheckingAccount.Withdraw(4200)
	fmt.Println("Balance after withdraw: ", firstCheckingAccount.GetBalance())

	firstCheckingAccount.Transfer(100_000., &secondCheckingAccount)

	fmt.Println("First checking account balance after transfer: ", firstCheckingAccount.GetBalance())
	fmt.Println("Second checking account balance after transfer: ", secondCheckingAccount.GetBalance())

	firstSavingsAccount := accounts.SavingsAccounts{
		Holder: customers.Holder{
			Name:       "Weslley",
			CPF:        "424.424.424-42",
			Ocuppation: "Quantitative Developer",
		},
		AgencyNumber:  042,
		AccountNumber: 424242,
		Operation:     42,
	}
	fmt.Println(firstSavingsAccount)

	// The interface make this possible. We can use a generic method.
	payBill(&firstCheckingAccount, 500)
	payBill(&firstSavingsAccount, 2000)
}
