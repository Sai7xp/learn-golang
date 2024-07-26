/*
	Another Example for Open Closed Principle
*/

package main

import (
	"fmt"
	"log"
)

type PaymentMethod interface {
	ConfirmPay()
}

type Wallet struct {
	amount   float64
	maxLimit int
}

func (w Wallet) ConfirmPay() {
	fmt.Printf("Paid %.2f using Wallet", w.amount)
}

type CreditCard struct {
	amount    float64
	isEnabled bool
}

func (w CreditCard) ConfirmPay() {
	if w.isEnabled {
		fmt.Printf("Paid %.2f using Credit Card", w.amount)
	} else {
		log.Fatalln("Credit card is not enabled!!")
	}
}

type Payment struct{}

func (p Payment) Checkout(pm PaymentMethod) {
	pm.ConfirmPay()
}

// func (cc CreditCard) Pay() {
// 	fmt.Printf("Paid %.2f using CreditCard", cc.amount)
// }

func openClosedYetAnotherExample() {
	p := Payment{}
	w := Wallet{amount: 12.98, maxLimit: 500.0}
	p.Checkout(w)

	// then in main()
	cc := CreditCard{amount: 22.33, isEnabled: true}
	p.Checkout(cc)

	pp := PayPal{amount: 45}
	p.Checkout(pp)
}

type PayPal struct {
	amount float64
}

func (pp PayPal) ConfirmPay() {
	fmt.Printf("Paid %.2f using PayPal", pp.amount)
}
