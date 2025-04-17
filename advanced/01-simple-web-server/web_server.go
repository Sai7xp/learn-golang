package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

const (
	PORT = ":5003"
)

type NewYearMessage struct {
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
	log.Println("Req received from : ", req.RemoteAddr)
	w.Write([]byte("bar"))
}

// Handler function for the POST request
func sendMessageHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Only POST method is allowed", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	// Decode JSON data from request body
	var msg NewYearMessage
	if err := json.NewDecoder(r.Body).Decode(&msg); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	log.Println("Received :", msg)

	// send JSON encoding of NewYearMessage{} in response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(msg)

	// observe the difference between Encode() vs w.Write()
	helloBytes := []byte(`hello`)
	json.NewEncoder(w).Encode(helloBytes) // Json Encode will encode []byte into base64 format
	w.Write(helloBytes)

}

func getMusicHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "only GET method is allowed", http.StatusBadRequest)
		return
	}

	// we can send image as "Content-Type" : "image/png"
	w.Header().Set("Content-Type", "audio/mpeg")

	file, err := os.Open("Arunachala.mp3")
	if err != nil {
		http.Error(w, "File not found", http.StatusNotFound)
		return
	}
	defer file.Close()

	// way 1
	http.ServeFile(w, r, "Arunachala.mp3")

	// way 2 of sending music file
	// _, err = io.Copy(w, file)
	// if err != nil {
	// 	http.Error(w, "Failed to stream audio", http.StatusInternalServerError)
	// }
}

func uploadFileHandler(w http.ResponseWriter, r *http.Request) {
	// r.ParseMultipartForm()
	file, header, err := r.FormFile("Profile")
	if err != nil {
		http.Error(w, "Error retrieving the file with name 'Profile'", http.StatusBadRequest)
		return
	}
	defer file.Close()

	// create a new local file
	newLocalFile, err := os.Create(header.Filename)
	if err != nil {
		http.Error(w, fmt.Sprintf("Unable to create the file : %v", err), http.StatusInternalServerError)
		return
	}
	defer newLocalFile.Close()

	fileBytes, _ := io.ReadAll(file)
	newLocalFile.Write(fileBytes)
	// Alternatively: io.Copy(newLocalFile, file)

	// Return success response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "Image Uploaded Successfully"})
}

func main() {

	// register the handlers
	http.HandleFunc("/foo", getHandler)                 // GET
	http.HandleFunc("/sendMessage", sendMessageHandler) // POST
	http.HandleFunc("/music", getMusicHandler)          // GET
	http.HandleFunc("/upload", uploadFileHandler)       // POST

	go func() {
		time.Sleep(time.Second) // wait for a sec until http server starts listening on given PORT
		makePostRequest()
	}()

	log.Printf("listening on port %s\n", PORT)
	if err := http.ListenAndServe(PORT, nil); err != nil {
		log.Fatal(err)
	}
}

// let's test the above post api /sendMessage
func makePostRequest() {
	fmt.Println("----> Making a POST request in Go <-------")
	var url = fmt.Sprintf("http://localhost%s/sendMessage", PORT)

	// dummy json payload
	reqBody := strings.NewReader(`
	{
		"currentYear" : 2025,
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
