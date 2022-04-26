package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func main() {
	// fileServer := http.FileServer(http.Dir("./examples"))
	// http.Handle("/", fileServer)
	http.HandleFunc("/setFile", ReceiveFile)
	http.HandleFunc("/getFile", SendFile)

	newpath := filepath.Join(".", "files")
	err := os.MkdirAll(newpath, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
