package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io"
	"log"
	"net/http"
	"strings"
	"sync"
)

// Struct to represent the JSON data
type Data struct {
	Message string `json:"message"`
	Year    int    `json:"currentYear"`
}

// Handler function for the POST request
func postRoute(w http.ResponseWriter, r *http.Request) {
	// Decode JSON data from request body
	var data Data
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Print received data
	log.Println("Received:", data)

	// Encode JSON data and send it back in the response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func main() {
	// Create a new router instance
	r := mux.NewRouter()

	var wg sync.WaitGroup

	// Define the route for the POST request
	r.HandleFunc("/testPost", postRoute).Methods("POST")

	wg.Add(1)
	go func() {
		defer wg.Done()
		// Start the HTTP server inside a go routine
		if err := http.ListenAndServe(":5000", r); err != nil {
			log.Fatal(err)
		}
	}()
	log.Println("Server started on port 5000")
	makePostRequest()
	wg.Wait()
}

// let's test the above post api /testPost
func makePostRequest() {
	fmt.Println("----> Making a POST request in Go <-------")
	const url = "http://localhost:5000/testPost"

	// dummy json payload
	reqBody := strings.NewReader(`
	{
		"currentYear" : 2024,
		"message": "Go is Awesome!😎"
	}
	`)
	res, err := http.Post(url, "application/json", reqBody)
	if err == nil {
		/// got success response
		fmt.Printf("%s status code %d\n", url, res.StatusCode)
	} else {
		/// something went wrong check the url
		log.Fatal(err)
	}
	defer res.Body.Close()

	content, _ := io.ReadAll(res.Body)
	fmt.Println("Response from /testPost API : ", string(content))
}
