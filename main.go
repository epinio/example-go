package main

import (
	"embed"
	"log"
	"net/http"
	"os"
)

//go:embed static/*
var staticFS embed.FS

func main() {
	http.Handle("/", http.FileServer(http.FS(staticFS)))
	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), nil))
}
