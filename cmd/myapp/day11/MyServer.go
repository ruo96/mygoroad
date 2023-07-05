package main

import (
	"fmt"
	"net/http"
	"os"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
	hostname, _ := os.Hostname()
	fmt.Fprintf(w, "Hello, you are visiting host: %s", hostname)
}

func main() {
	http.HandleFunc("/hello", indexHandler)
	fmt.Println("Server starting on port: 8088...")
	http.ListenAndServe(":8088", nil)
}
