/*
* Created on 26 Feb 2024
* @author Sai Sumanth
 */

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// /
// / Modules in GO
// /
// / Let's see how to import third party go modules
// /
func main() {
	fmt.Println("Modules in Go")
	/*
		Adding mux module to out module
		command : go get -u github.com/gorilla/mux
		Mux is a ackage gorilla/mux implements a request router and dispatcher for
		matching incoming requests to their respective handler.
		The name mux stands for "HTTP request multiplexer".

		Even after adding mux we won't be able to see any other files imported to our module folder
		just like node_modules. that's because files will be stored at GOPATH

		Mux files can be found at /Users/webknottechnologies/go/pkg/mod/cache/download/github.com/gorilla


		go mod verify -> to verify the packages

		go list -> will list module name // github.com/sai7xp/learn-golang

		go list all -> lists all the module names

		go mod graph -> displays the dependency graph

		go mod vendor

	*/

	/// let create a router using mux
	r := mux.NewRouter()

	r.HandleFunc("/", HomeHandler).Methods("GET")
	r.HandleFunc("/settings/{userId}", SettingsHandler).Methods(http.MethodPost)
	fmt.Println("Listening at port 5000")
	if err := http.ListenAndServe(":5000", r); err != nil {
		/// then there's some issue with creating the server
		log.Fatal(err)
	}

}

func HomeHandler(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("<h3>Welcome Home. Hit POST /settings/{userId} to see your details</h3>"))

}

// post api /settings/{userId} handler
func SettingsHandler(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)

	var jsonD = make(map[string]interface{})
	jsonD["message"] = "Welcome to Your Settings Page"
	jsonD["name"] = fmt.Sprintf("Hi %s", vars["userId"])
	if userIdAsInteger, err := strconv.ParseInt(vars["userId"], 10, 64); err == nil {
		jsonD["userId"] = userIdAsInteger
	} else {
		jsonD["userId"] = vars["userId"]
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	// if responseDataToUser, err := json.Marshal(jsonD); err == nil {
	json.NewEncoder(w).Encode(jsonD)
	// fmt.Fprintf(w, "Server 5000 is up and Running")
	// } else {
	// 	fmt.Fprintf(w, "OOPS %v Details not Founc\n", vars["userId"])
	// 	w.WriteHeader(http.StatusBadRequest)
	// }
}
