package main

import (
	"fmt"
	"mygo/cmd/myapp/day2"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello, world!")
}

func main() {
	//http.HandleFunc("/", helloHandler)
	//log.Fatal(http.ListenAndServe(":8088", nil))

	day2.CalIota()
}
