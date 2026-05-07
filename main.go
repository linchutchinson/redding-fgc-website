package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	static_fs := os.DirFS("static")
	static_fs_handler := http.FileServer(http.FS(static_fs))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		static_fs_handler.ServeHTTP(w, r)
	})

	fmt.Println("Hello RFGC")
	http.ListenAndServe(":8080", nil)
}