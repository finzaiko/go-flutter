// package main

// import (
// 	"net/http"
// )

// func main() {
// 	fs := http.FileServer(http.Dir("public"))
// 	http.Handle("/", fs)
// 	http.ListenAndServe(":8080", nil)
// }

package main

import (
	"embed"
	"io/fs"
	"log"
	"net/http"
)

//go:embed public/*
var static embed.FS

func main() {
	contentStatic, _ := fs.Sub(static, "public")
	http.Handle("/", http.FileServer(http.FS(contentStatic)))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
