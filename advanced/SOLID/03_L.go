/*
* Created on 01 April 2024
* @author Sai Sumanth
 */
package main

import (
	"fmt"
)

func Liskov() {
	fmt.Println("\n3️⃣ Liskov Substitution Principle - GoLang SOLID Principles")
	// Object of a superclass shall be replaceble with objects of its subclasses
	// without breaking the application.
	e := Eagle{}
	p := Penguin{}

	travel(&e)
	travel(&p) // ❌ Penguins can't fly, so this fn call will panic

	LiskovImplementation()
}

type Bird interface {
	Fly()
}

// According to the Liskov principle, objects of a superclass should be replaceable with
// objects of its subclasses.
// However, when the [travel] method is called and a penguin instance is passed,
// the application will break since penguins can't fly, violating the Liskov principle
// ❌ Liskov Principle Violation
func travel(b Bird) {
	b.Fly()
}

type Eagle struct{}

func (e *Eagle) Fly() {
	fmt.Println("Bird is flying")
}

type Penguin struct{}

func (p *Penguin) Fly() {
	fmt.Println("This bird can't fly")
	// 🚨
	panic(1)
}

// ✅

type FlyingBirds interface {
	TravelByAir()
}

// flying birds can send letters
func sendLetters(fb FlyingBirds) {
	fb.TravelByAir()
}

type NonFlyingBirds interface {
	TravelByWalk()
}

func sendSlowLetters(nfb NonFlyingBirds) {
	nfb.TravelByWalk()
}

// flying bird
type Pegion struct{}
type Crow struct{}

func (p *Pegion) TravelByAir() {
	fmt.Println("Pegion Started Travelling by Air 🚁")
}
func (cr *Crow) TravelByAir() {
	fmt.Println("Crow Started Travelling by Air 🚁")
}

// non flying bird
type Kiwi struct{}

func (ki *Kiwi) TravelByWalk() {
	fmt.Println("Kiwi Started going by Walk 🚶")
}

func LiskovImplementation() {
	// Non Flying Birds
	k := Kiwi{}

	// Flying Birds
	pe := Pegion{}
	cr := Crow{}

	/// Letters can be sent through Pegion or Crow
	sendLetters(&pe)
	sendLetters(&cr)

	// sendLetters(&k) ❌ Kiwi bird(non flying) can't be used to send letters

	sendSlowLetters(&k)
}
