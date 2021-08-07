package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
)

// The flag package provides a default help printer via -h switch
var versionFlag *bool = flag.Bool("v", false, "print the version number.")

var podName = ""

type EchoResponse struct {
	Message string `json:"message"`
	PodName string `json:"podname"`
}

func echo(w http.ResponseWriter, r *http.Request) {
	response := EchoResponse{
		Message: "echo",
		PodName: podName,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
	log.Printf("%s: /echo %d", podName, http.StatusOK)
}

func main() {
	// scan arguments
	flag.Parse()

	// display application version
	if *versionFlag {
		fmt.Printf("version: %v (build %v by %v on %v at %v)\n", BuildVersion, BuildCommit, BuildUser, BuildDate, BuildTime)
		os.Exit(0)
	}

	podName, _ = os.LookupEnv("POD_NAME")

	mux := http.NewServeMux()
	mux.HandleFunc("/echo", echo)
	log.Printf("starting echo service on %s:%d", podName, 8080)
	log.Fatal(http.ListenAndServe(":8080", mux))
}
