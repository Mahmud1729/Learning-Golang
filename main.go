package main

import (
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello world")
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "I am Khatami. I am an Undergraduate Student")
}
func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/hello", helloHandler)
	mux.HandleFunc("/about", aboutHandler)

	fmt.Println("Server running on port 3000")

	err := http.ListenAndServe(":3000", mux)
	if err != nil {
		fmt.Println("There is an error running this server: ", err)
	}
}
