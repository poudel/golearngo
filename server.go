package main

import (
	"fmt"
	"net/http"
	"log"
)


func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world! %s", r.URL.Path)
}


func main () {
	http.HandleFunc("/", handler)

	fmt.Println("Listening at http://localhost:8799")
	log.Fatal(http.ListenAndServe(":8799", nil))
}
