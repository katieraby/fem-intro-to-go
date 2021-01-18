package main

import (
	"fmt"
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Home!")
}

func main() {
	http.HandleFunc("/", home)
	fmt.Println("Building a web server")
	// even wrapped in log fatal, the code will still execute
	// will fallback to log fatal if there is an error
	log.Fatal(http.ListenAndServe(":8080", nil))
}
