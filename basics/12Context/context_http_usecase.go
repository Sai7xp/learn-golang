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

func contextWithHTTP() {
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
	result := make(chan int)

	go func() {
		computationResult, err := expensiveComputation(ctx, 47)
		if err != nil {
			log.Println("Computation error:", err)
			return
		}
		fmt.Println("COMPUTATION RESULT: ", computationResult)
		result <- computationResult
	}()

	select {
	case <-ctx.Done():
		log.Println("client canceled the request: ", ctx.Err())
		return
	case res := <-result:
		log.Println("Data fetched")
		w.Write([]byte(fmt.Sprintf("Calculation Result: %d", res)))
		return
	}
}

func expensiveComputation(ctx context.Context, n int) (int, error) {
	if n <= 1 {
		return n, nil
	}

	// Check context cancellation periodically.
	select {
	case <-ctx.Done():
		// whenever the client is disconnected we should stop the computation as well
		// otherwise it's waste of resources
		return 0, ctx.Err()
	default:
		// Continue computation
	}

	a, err := expensiveComputation(ctx, n-1)
	if err != nil {
		return 0, err
	}
	b, err := expensiveComputation(ctx, n-2)
	if err != nil {
		return 0, err
	}
	return a + b, nil
}
