/*
* Created on  12 March 2024
* @author Sai Sumanth
 */
package main

import (
	"fmt"
)

/*
* Embedding example
 */

// User struct represents basic user in our app
type User struct {
	Username string
	email    string
}

// display user info
func (u User) displayUserInfo() {
	fmt.Printf("User: %s, Email: %s\n", u.Username, u.email)
}

type Admin struct {
	User        // Embedding User Struct (Admin 'Has-A' User) This is not true inheritance
	Permissions []string
}

// Method specific to Admin for displaying permissions.
func (a Admin) DisplayPermissions() {
	/*
		we can directly access Username as `a.Username` instead of `a.user.Username`
	*/
	fmt.Printf("Admin %s has the following permissions: %v\n", a.Username, a.Permissions)
}

// pass Admin to this function and see if it works, It won't work because go doesn't support true inheritance
func PrintUser(u User) {
	fmt.Println(u)
}

func RunEmbeddingUsage() {
	user := User{
		email:    "user@go.dev",
		Username: "Go Dev",
	}
	user.displayUserInfo()
	admin := Admin{
		User: User{
			email:    "admin@go.dev",
			Username: "Go Supreme",
		},
		Permissions: []string{"write", "read"},
	}

	// var myUser User = Admin{} // ❌ this is not allowed in Go, but in Java we can do this. assigning child object to parent class
	// PrintUser(admin) // ❌

	admin.displayUserInfo() // 😎 No need to access displayUserInfo() like `a.user.displayUserInfo()`
	admin.DisplayPermissions()

	// check the below example as well
	// https://gobyexample.com/struct-embedding

	type Guest struct {
		*User // this has to be initialized since it is a pointer, otherwise it will panic when we access methods belongs to User
	}
	guest := Guest{}
	fmt.Println("guest:", guest)
	// guest.displayUserInfo() // 🚨 panic - nil pointer dereference. since *User is embedded in Guest

	// Interface Embedding
	RunInterfaceEmbedding()
}

/*
Can we embed interface into structs ?
Yes, but it is not recommended. But we can embed interface into another interface
`io.ReadWriter` interface embeds Reader and Writer interfaces
*/

type Service interface {
	GetUser(userId string) string
	SayHello()
}

type MyMock struct {
	// this is an interface how will you initialize this field while creating object for MyMock ?
	// we have to create one struct which implements Service and use it here. then what's the point of MyMock
	Service // interface embedded in a struct, which is not recommended
}

func (m MyMock) SayHello() {
	fmt.Println("SayHello method belongs to MyMock struct")
}

func RunInterfaceEmbedding() {
	mm := MyMock{}
	mm.SayHello()   // if we don't define this method for MyMock, then this will also cause panic
	mm.GetUser("1") // 🚨 panic: GetUser behaviour is not implemented, it's just an interface method
}
