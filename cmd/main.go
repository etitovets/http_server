package main

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

var message string

func getRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got / request\n")
	if _, err := io.WriteString(w, message); err != nil {
		fmt.Printf("error writing response: %s\n", err.Error())
	}
}

func getHello(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got /hello request\n")
	if _, err := io.WriteString(w, "Hello, HTTP!\n"); err != nil {
		fmt.Printf("error writing response: %s\n", err.Error())
	}
}

func main() {
	flag.StringVar(&message, "m", "hello", "message")
	flag.Parse()

	f, err := os.ReadFile("/cmd/config")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Loaded config: %s\n", f)

	http.HandleFunc("/", getRoot)
	http.HandleFunc("/hello", getHello)

	err = http.ListenAndServe(":8081", nil)

	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}
}
