package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Redding FGC Landing Page")
	})

	fmt.Println("Hello RFGC")
	http.ListenAndServe(":8080", nil)
}