package main

import (
	"encoding/json"
	"fmt"
	"learngo/advanced/02-go-mux-pro/datamodels"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gorilla/mux"
)

// our cats dummy database
var allCatsDb = []datamodels.Cat{}

func main() {
	fmt.Println("Welcome to Go Mux API")

	// dummyCourse := datamodels.Cat{}
	/// add two cats to store
	allCatsDb = append(allCatsDb, datamodels.Cat{Name: "Daalee",
		AgeInMonths: 2, WeightInKgs: 45, Color: "Orange", IsMale: true,
		DateOfBirth: time.Now(), Vaccinations: []string{"vacc1"},
		Owner: &datamodels.Owner{Name: "sai", IsMale: true}})

	// create router
	router := mux.NewRouter()

	router.HandleFunc("/", serveCatsHome).Methods("GET")

	/// GET : return all cats info
	router.HandleFunc("/getCats", getAllCats).Methods(http.MethodGet)

	/// GET : grab `catName` from params and return specific cat details
	router.HandleFunc("/getCatDetails/{catName}", getCatDetails).Methods(http.MethodGet)

	/// POST : Add New Cat
	router.HandleFunc("/addCat", addCatHandler).Methods("POST")

	/// PUT : Update existing Cat Details
	router.HandleFunc("/update", updateCatHandler).Methods("PUT")

	/// DELETE: Remove Cat
	router.HandleFunc("/deleteCat/{catName}", deleteCatHandler).Methods("DELETE")

	/// wrong route handler
	router.NotFoundHandler = http.HandlerFunc(routeNotFoundHandler)

	fmt.Println("Server started at port 5050")
	log.Fatal(http.ListenAndServe(":5050", router))

}

/// Controllers/Handlers

// serve Cats Home
// http.ResponseWriter to write the response back to the client,
// and the *http.Request to access information about the incoming request if needed
func serveCatsHome(w http.ResponseWriter, r *http.Request) {
	// 💡
	// w is an interface and interfaces are passed by value,
	// but they're lightweight and efficient to copy.

	// 💡
	// r is a struct and structs are typically passed by pointer
	// to avoid the overhead of copying large amounts of data.
	fmt.Println("User Hit Home Page")
	w.Write([]byte("Welcome to Cats Home"))
}

// fetch all Cat details
func getAllCats(w http.ResponseWriter, r *http.Request) {
	fmt.Println("User Triggered Get All Cats Info")
	w.Header().Set("Content-Type", "application/json")

	responseMap := make(map[string]interface{})
	responseMap["data"] = allCatsDb
	responseMap["message"] = "Cats Data Fetched Successfully!😻"
	json.NewEncoder(w).Encode(responseMap)
}

// fetch specific cat details
func getCatDetails(w http.ResponseWriter, r *http.Request) {
	fmt.Println("User tried to fetch the specific cat details")
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	/// loop through all the cats and find the one user wants
	for _, v := range allCatsDb {
		if strings.EqualFold(v.Name, params["catName"]) {
			json.NewEncoder(w).Encode(v)
			return
		}
	}
	json.NewEncoder(w).Encode(map[string]string{"message": "No Cat Details Found"})
}

// for adding a Cat 😻
func addCatHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("User tried to add New Cat")
	w.Header().Set("Content-Type", "application/json")

	/// what if body is empty
	if r.Body == nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Please send cat details")
		return
	} else {
		fmt.Println(r.Body)
	}

	var newCat datamodels.Cat
	json.NewDecoder(r.Body).Decode(&newCat)
	if newCat.Owner == nil {
		// expected := `{"alive": true}`
		// json.NewEncoder(w).Encode(`{"message":"Owner details are required to Add Cat Details"}`)
		w.Write([]byte(`{"message":"Owner details are required to Add Cat Details"}`))
		return
	} else {
		allCatsDb = append(allCatsDb, newCat)
		message := fmt.Sprintf("%s New Cat Added Successfully!", newCat.Name)
		newCat.PrintInfo()
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(map[string]string{"message": message, "data": newCat.ToString()})
	}
}

// TODO: update Cat handler
func updateCatHandler(w http.ResponseWriter, r *http.Request) {
}

// Delete Cat Handler
func deleteCatHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	/// get the catname from params
	for index, eachCat := range allCatsDb {
		if strings.EqualFold(eachCat.Name, params["catName"]) {
			allCatsDb = append(allCatsDb[:index], allCatsDb[index+1:]...)
			w.WriteHeader(http.StatusOK)
			message := fmt.Sprintf("%s Deleted Successfully from db", params["catName"])
			json.NewEncoder(w).Encode(map[string]string{"message": message})
			return
		}
	}
	w.WriteHeader(http.StatusBadRequest)
	w.Write([]byte("Cat Details Not Found"))
}

func routeNotFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte(`{"message":"🐱 Meow! It seems like you've wandered off into uncharted territory. 
    Our feline friends couldn't find the page you're looking for. 🐾
    But don't worry, there are plenty of purrfect paths to explore on our Cats API! 😺"}`))
}
