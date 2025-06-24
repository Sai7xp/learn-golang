package main

import "fmt"

type BankAccount struct {
	name   string
	amount int
}

func (b *BankAccount) deposit(amount int) {
	b.amount += amount
	fmt.Println("Amount added in BankAccount", b.amount)
}

func (b *BankAccount) withdraw(amount int) {
	b.amount -= amount
	fmt.Println("Withdaw in BankAccount: ", amount)
}

type FDAccount struct {
	BankAccount
}

func (f *FDAccount) withdraw(amount int) {
	panic("Withdrawl not supported in FD account")
}

// refer to Liskov Java example. this code is just to understand how Go doesn't support true inheritance even if we use composition
// liskov principle talks more about superclass and subclasses. Interface seggregation principle talks about interfaces more
func RunLiskovExample2() {
	acc := &BankAccount{name: "uncle bob", amount: 1000}
	acc.deposit(99976)
	acc.withdraw(12)

	fdAcc := &FDAccount{BankAccount{name: "jane", amount: 7990}}
	fdAcc.withdraw(89) // 🔴 panic, withdraw not supported for

	var myAcc BankAccount
	// myAcc = FDAccount{}
	// this is not allowed in Go, even though FDAccount kind of extends BankAccount.
	// But we can do this in Java because it supports true inheritance, so we can substitute subclass in the place of parent class

	fmt.Println(myAcc)
}
