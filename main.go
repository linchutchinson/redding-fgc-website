package main

import (
	"fmt"
	"net/http"
	"embed"
)

//go:embed static
var static embed.FS

func main() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.FS(static))))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "static/index.html")
	})

	fmt.Println("Hello RFGC")
	http.ListenAndServe(":8080", nil)
}