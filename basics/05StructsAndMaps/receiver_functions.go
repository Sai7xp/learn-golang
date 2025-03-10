/*
* Created on 29 Feb 2024
* @author Sai Sumanth
 */
package main

import (
	"errors"
	"fmt"
)

/*
** Receiver Functions **
- Methods/Functions that are associated with a specific type
- They are useful for encapsulating logic that operates on the data of a particular type

-- Types of Receivers --
1. Value Receiver - operates on a copy of type t
2. Pointer Receiver - operates on the actual instance of type t
*/

func receiverFunctions() {
	onEmailType()

	onStruct()
}

type Email string

func (s Email) IsEmailValid() error {
	if len(s) < 5 {
		return errors.New("invalid email")
	}
	return nil
}

func onEmailType() {
	fmt.Println("Receiver Functions in GO")
	userEmail := Email("sai@dev.ai")
	if err := userEmail.IsEmailValid(); err != nil {
		panic(err)
	} else {
		fmt.Println(userEmail, " is Valid")
	}
}

func onStruct() {
	account1 := &BankAccount{ownerName: "Ram", balance: 0}
	account1.Deposit(500)
	account1.CheckBalance()
	if err := account1.Withdraw(200); err == nil {
		fmt.Println("Amount Withdrawl success. Remaining balance : ", account1.balance)
	} else {
		fmt.Println(err)
	}
}

type BankAccount struct {
	ownerName string
	balance   float64
}

// Value Receiver on BankAccount to check balance
func (b BankAccount) CheckBalance() float64 {
	// Operates on a copy of b
	fmt.Println("Your account balance is: ", b.balance)
	return b.balance
}

// Pointer Receiver to deposit money
// If we use a value receiver here, changes to balance will not reflect
func (b *BankAccount) Deposit(amount float64) {
	b.balance += amount
	fmt.Printf("%.2f has been added to your account balance. Total Balance: %.2f\n", amount, b.balance)
}

func (b *BankAccount) Withdraw(amount float64) error {
	if b.balance < amount {
		return errors.New("insufficient balance")
	}
	b.balance -= amount
	return nil

}

/*
Method Chaining : calling multiple methods on the same object in a single statement.

	import sq "github.com/Masterminds/squirrel"
	users := sq.Select("*").From("users").Join("emails USING (email_id)")

We can achieve this by having each method return the receiver (the object itself),
allowing the next method to be called on the returned value.

*/
