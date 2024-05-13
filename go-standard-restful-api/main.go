package main

import (
	"fmt"
	"net/http"
	"time"

	"novisshrivastava.stdrestapi.com/routes"
)

func main() {
	// Register routes
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, World!")
	})
	http.HandleFunc("/cpu-intensive", func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		sum := 0
		for i := 0; i < 100000000; i++ {
			sum += i
		}
		elapsed := time.Since(start)
		fmt.Fprintf(w, "CPU-intensive task completed in %s\n", elapsed)
	})
	http.HandleFunc("/posts", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Inside 1")
		switch r.Method {
		case http.MethodGet:
			fmt.Println("Inside 2")
			routes.PostsResource{}.List(w, r)
		case http.MethodPost:
			routes.PostsResource{}.Create(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	// Start the server
	http.ListenAndServe(":8080", nil)
}
