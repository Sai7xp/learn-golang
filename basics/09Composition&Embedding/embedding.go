/*
* Created on  12 March 2024
* @author Sai Sumanth
 */
package main

import (
	"fmt"
)

/*
* Embedding Real-life example
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

func main() {
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

	/*
		More Examples
	*/
	RunComicExample()

	// check the below example as well
	// https://gobyexample.com/struct-embedding
}

/*
Another Example for Embedding
*/

type Comic struct {
	Universe string
}

func (comic *Comic) prinUniverse() {
	fmt.Println("Comic Universe")
}

type Marvel struct {
	// anonymous field,
	// the struct is embedded
	Comic
	Owner string
}

func (comic *Marvel) prinUniverse() {
	fmt.Println("Marvel Universe")
}

type DC struct {
	Comic
	Owner string
}

func RunComicExample() {
	fmt.Println("-------------- Another Comic Example --------------")
	cm := Comic{"Comic MilkyWay"}
	cm.prinUniverse()

	mar := Marvel{Comic{Universe: "MCU"}, "ironman"}
	dc := DC{Comic{Universe: "DCU"}, "batman"}
	fmt.Println(mar.Universe)
	mar.prinUniverse() // Marvel Universe - since we have receiver fn for Marvel it will print Marvel Universe
	// prinUniverse fn of Marvel will override the prinUniverse of Comic

	fmt.Println(dc.Universe)
	dc.prinUniverse() // since we don't have any receiver function for DC - Comic Universe will be printed

}
