package main

import (
	"fmt"
	"net/http"
)

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprint(w, "<h1>Welcome to my awesome site!</h1>")
	if err != nil {
		return
	}
}

func main() {
	http.HandleFunc("/", handlerFunc)
	fmt.Println("Starting the server on :3000...")
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		return
	}
}
