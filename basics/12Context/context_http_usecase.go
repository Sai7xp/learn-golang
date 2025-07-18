package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
)

const (
	PORT = ":7003"
)

// Open the URL in browser - http://localhost:7003/get and close the tab
func contextWithHTTPServer() {
	http.HandleFunc("/get", getHandler)
	httpServer := http.Server{
		Addr: PORT,
	}
	log.Printf("Starting server on %s\n", PORT)
	if err := httpServer.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}

func getHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	log.Println("processing request...")
	computationResult, err := expensiveComputation(ctx, 60)
	if err != nil {
		log.Println("Computation error:", err)
		return
	}
	fmt.Println("COMPUTATION RESULT: ", computationResult)
	w.Write([]byte(fmt.Sprintf("Calculation Result: %d", computationResult)))
}

// nth fibonacci
func expensiveComputation(ctx context.Context, n int) (int, error) {

	// Check context cancellation periodically.
	select {
	case <-ctx.Done():
		fmt.Println("CLIENT DISCONNECTED. ")
		// whenever the client is disconnected we should stop the computation as well
		// otherwise it's waste of resources
		return 0, ctx.Err()
	default:
		if n <= 1 {
			return n, nil
		}
		a, err := expensiveComputation(ctx, n-1)
		if err != nil {
			return 0, err
		}
		b, err := expensiveComputation(ctx, n-2)
		if err != nil {
			return 0, err
		}
		// return (n-1) + (n-2) // fibonacci
		return a + b, nil
	}

}
