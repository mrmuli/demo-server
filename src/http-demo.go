package main

import (
	"fmt"
	"github.com/Unleash/unleash-client-go"
	"net/http"
)

func init() {
	unleash.Initialize(
		unleash.WithListener(&unleash.DebugListener{}),
		unleash.WithAppName("http-demo"),
		unleash.WithUrl("http://localhost:4242/api/"),
		unleash.WithCustomHeaders(http.Header{"Authorization": {"very-secret-token"}}),
	)
}

func handler(w http.ResponseWriter, r *http.Request) {
	if unleash.IsEnabled("http-demo.hello") {
		fmt.Fprintf(w, "Congratulations! this feature is enabled\n")
	} else {
		fmt.Fprintf(w, "Hello this feature isn't there yet!\n")
	}
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":3030", nil)
}
