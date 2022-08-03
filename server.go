package main

import (
	"log"
	"net/http"
)

func serve() {
	log.Println("starting http dev server on :3000")
	log.Fatal(http.ListenAndServe(":3000", http.FileServer(http.Dir("."))))
}
