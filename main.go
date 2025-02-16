package main

import (
	"embed"
	"fmt"
	"log"
	"net/http"
)

//go:embed index.html
var indexHTML embed.FS

func main() {
	addr := ":8000"

	http.Handle("/", http.FileServer(http.FS(indexHTML)))

	http.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("OK"))
	})

	fmt.Println("listening at:", addr)
	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatal(err)
	}
}
