package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
)

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

	w.WriteHeader(status_code)

	fmt.Fprintf(w, message)
}

func main() {
	http.HandleFunc("/status/", mirrorStatus)

	fmt.Println("Listening at 8799")
	log.Fatal(http.ListenAndServe(":8799", nil))
}
