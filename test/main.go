package main

import (
	"embed"
	"net/http"
)

//go:embed index.html
var index embed.FS

func main() {
	http.Handle("/", http.FileServer(http.FS(index)))
	http.ListenAndServe(":8080", nil)
}
