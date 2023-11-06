package main

import (
	"fmt"
	"net/http"

	"github.com/a-h/templ"
)

func main() {
	component := hello("john")

	http.Handle("/", templ.Handler(component))
	fmt.Println("Listening on :3001")
	http.ListenAndServe(":3001", nil)
}
