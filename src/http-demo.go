package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Congratulations! Your Go application is running!.")
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":3030", nil)
}
