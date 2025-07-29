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
	User        // Embedding User Struct (clearly states that Admin "is-a" User relationship)
	Permissions []string
}

// Method specific to Admin for displaying permissions.
func (a Admin) DisplayPermissions() {
	/*
		we can directly access Username as `a.Username` instead of `a.user.Username`
	*/
	fmt.Printf("Admin %s has the following permissions: %v\n", a.Username, a.Permissions)
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

	admin.displayUserInfo() // 😎 No need to access displayUserInfo() like `a.user.displayUserInfo()`
	admin.DisplayPermissions()

	// check the below example as well
	// https://gobyexample.com/struct-embedding

	guest := Guest{}
	fmt.Println("guest:", guest)
	// guest.displayUserInfo() // 🚨 panic - nil pointer dereference. since *User is embedded in Guest

	// Interface Embedding
	RunInterfaceEmbedding()
}

type Guest struct {
	*User
}

/*
Can we embed interfaces into structs ?
Yes, but it is not recommended. We can embed interface into another interface
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
