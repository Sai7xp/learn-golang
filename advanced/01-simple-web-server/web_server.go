package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"time"
)

const (
	PORT = ":5003"
)

type Data struct {
	Message string `json:"message"`
	Year    int    `json:"currentYear"`
}

/*
This get handler can be triggered from terminal as well, instead of using postman.
connect to the server using `nc localhost <PORT>`
then start sending messages to http server

GET /foo HTTP/1.1
Host: localhost

Send the above message as raw text, server will respond back
*/
func getHandler(w http.ResponseWriter, req *http.Request) {
	log.Println("Req from : ", req.RemoteAddr)
	w.Write([]byte("bar"))
}

// Handler function for the POST request
func postRoute(w http.ResponseWriter, r *http.Request) {
	// Decode JSON data from request body
	var data Data
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	log.Println("Received:", data)

	// Encode JSON data and send it back in the response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func main() {

	// register the handlers
	http.HandleFunc("/testPost", postRoute)
	http.HandleFunc("/foo", getHandler)

	go func() {
		time.Sleep(time.Second) // wait for a sec until http server starts listening on given PORT
		makePostRequest()
	}()

	log.Printf("listening on port %s\n", PORT)
	if err := http.ListenAndServe(PORT, nil); err != nil {
		log.Fatal(err)
	}
}

// let's test the above post api /testPost
func makePostRequest() {
	fmt.Println("----> Making a POST request in Go <-------")
	var url = fmt.Sprintf("http://localhost%s/testPost", PORT)

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
