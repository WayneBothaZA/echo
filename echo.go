package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
)

// The flag package provides a default help printer via -h switch
var versionFlag *bool = flag.Bool("v", false, "print the version number.")
var hostName = ""
var portNumber = "8080"

// EchoRequest defines the JSON structure of the request message
type EchoRequest struct {
	Message string `json:"message"`
}

// EchoResponse defines the JSON structure of the response message
type EchoResponse struct {
	Message   string `json:"message"`
	Hostname  string `json:"hostname"`
	UserAgent string `json:"useragent"`
}

func health(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func echo(w http.ResponseWriter, r *http.Request) {
	var req EchoRequest
	err := decodeJSONBody(w, r, &req)
	if err != nil {
		var mr *malformedRequest
		if errors.As(err, &mr) {
			log.Printf("%s: /echo %d - %s", hostName, mr.status, mr.msg)
			http.Error(w, mr.msg, mr.status)
		} else {
			log.Printf("%s: /echo - %v %s (%s)", hostName, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError), err.Error())
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}
		return
	}

	response := EchoResponse{
		Message:   req.Message,
		Hostname:  hostName,
		UserAgent: r.Header.Get("User-Agent")}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
	log.Printf("%s: /echo %d", hostName, http.StatusOK)
}

func main() {
	// scan arguments
	flag.Parse()

	// display application version
	if *versionFlag {
		fmt.Printf("version: %v (build %v by %v on %v at %v)\n", BuildVersion, BuildCommit, BuildUser, BuildDate, BuildTime)
		os.Exit(0)
	}

	hostName, _ = os.LookupEnv("HOSTNAME")
	if envPort, found := os.LookupEnv("PORT"); found {
		portNumber = envPort
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/echo", echo)
	mux.HandleFunc("/health", health)

	url := ":" + portNumber
	log.Printf("starting echo service on %s%s", hostName, url)
	log.Fatal(http.ListenAndServe(url, mux))
}
