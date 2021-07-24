package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
)

// The flag package provides a default help printer via -h switch
var versionFlag *bool = flag.Bool("v", false, "print the version number.")

func echo(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("echo\n"))
	log.Printf("/echo %d", http.StatusOK)
}

func main() {
	// scan arguments
	flag.Parse()

	// display application version
	if *versionFlag {
		fmt.Printf("version: %v (build %v by %v on %v at %v)\n", BuildVersion, BuildCommit, BuildUser, BuildDate, BuildTime)
		os.Exit(0)
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/echo", echo)
	log.Printf("starting echo service on:%d", 8080)
	log.Fatal(http.ListenAndServe(":8080", mux))
}
