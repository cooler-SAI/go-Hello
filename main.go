package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Received request from %s for path %s", r.RemoteAddr, r.URL.Path)
	hostname, err := os.Hostname()
	if err != nil {
		log.Fatal("Can't get Hostname: ", err)
	}
	_, err2 := fmt.Fprintf(w, "Hello from Kubernetes! I'm serving from pod: %s\n", hostname)
	if err2 != nil {
		log.Fatal("Can't write response: ", err2)
		return
	}
}

func main() {

	fmt.Println("Hello World")

}
