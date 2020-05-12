package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
)

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Hello World!\n")
}

func getenv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return fallback
}

func main() {
	http.HandleFunc("/", hello)

	portStr := getenv("PORT", "3000")
	port, err := strconv.Atoi(portStr)
	if err != nil {
		log.Fatalf("Given an invalid port \"%s\"!\n", portStr)
	}

	log.Printf("Listening on port %d\n", port)
	http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}
