package main

import (
	"fmt"
	"io/fs"
	"net/http"
	"embed"
)

//go:embed static
var static embed.FS

func main() {
	static_fs, _ := fs.Sub(static, "static")
	static_fs_handler := http.FileServer(http.FS(static_fs))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		static_fs_handler.ServeHTTP(w, r)
	})

	fmt.Println("Hello RFGC")
	http.ListenAndServe(":8080", nil)
}