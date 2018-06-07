package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"encoding/json"
)

type StatusResponse struct {
	Message string `json:"message"`
	StatusCode int `json:"status_code"`
	Origin string `json:"origin"`
}


func mirrorStatus(w http.ResponseWriter, r *http.Request) {
	status := strings.TrimRight(r.URL.Path[len("/status/"):], "/")

	// try to see if it can be converted to integer
	status_code, err := strconv.Atoi(status)
	if err != nil {
		fmt.Fprintf(w, "Invalid status code: %s", status)
		return
	}

	message := http.StatusText(status_code)

	if len(message) == 0 {
		fmt.Fprintf(w, "Status code does not exist: %s", status)
		return
	}

	response := StatusResponse{
		message,
		status_code,
		r.RemoteAddr,
	}

	js, err := json.Marshal(response)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status_code)
	w.Write(js)
}


func mirrorIp(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, r.RemoteAddr)
}

func main() {
	http.HandleFunc("/status/", mirrorStatus)
	http.HandleFunc("/ip/", mirrorIp)

	fmt.Println("Listening at 8799")
	log.Fatal(http.ListenAndServe(":8799", nil))
}
