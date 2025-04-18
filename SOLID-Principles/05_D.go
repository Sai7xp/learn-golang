/*
* Created on 14 April 2024
* @author Sai Sumanth
 */
package main

import (
	"fmt"
)

func DependencyInversion() {
	fmt.Println("\n5️⃣ Dependency Inversion Principle - GoLang SOLID Principles")
	// checkout go-secure-pdf in golang-projects repository
	// while writing methods to encrypt/decrypt or lock/unlock the pdfs
	// if we directly depend on a specific package it becomes difficult in future to use another dependency.
	// so we have to depend on abstractions
}
