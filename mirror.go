package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type StatusResponse struct {
	Message    string `json:"message"`
	StatusCode int    `json:"status_code"`
	Ip         string `json:"ip"`
	Method string `json:"method"`
}

func cleanIp(addr string) string {
	occur := strings.LastIndex(addr, ":")
	return addr[:occur]
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
		cleanIp(r.RemoteAddr),
		r.Method,
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

type IpResponse struct {
	Ip string `json:"ip"`
}

func mirrorIp(w http.ResponseWriter, r *http.Request) {
	response := IpResponse{cleanIp(r.RemoteAddr)}
	js, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func main() {
	http.HandleFunc("/status/", mirrorStatus)
	http.HandleFunc("/ip/", mirrorIp)

	fmt.Println("Listening at 8799")
	log.Fatal(http.ListenAndServe(":8799", nil))
}
