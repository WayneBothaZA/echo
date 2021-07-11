package main

import (
	"log"
	"net/http"
)

func echo(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("echo\n"))
	log.Printf("/echo %d", http.StatusOK)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/echo", echo)
	log.Printf("starting echo service on:%d", 8080)
	log.Fatal(http.ListenAndServe(":8080", mux))
}
