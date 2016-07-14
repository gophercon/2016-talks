// Command web is a trivial web server.
package main

import (
	"fmt"
	"log"
	"net/http"
)

type Celsius float64

func (c Celsius) String() string {
	return fmt.Sprintf("%g°C", c)
}

const balmy Celsius = (80 /*°F*/ - 32.0) * 5 / 9

func handleRoot(w http.ResponseWriter, req *http.Request) {
	log.Print("Request from %s", req.URL)
	w.Header().Set("Content-Type", "text/plain")
	fmt.Fprintf(w, "The temperature is %v", balmy)
}

func main() {
	http.HandleFunc("/", handleRoot)

	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatal(err)
	}
}

type handler struct{}

func (handler) ServeHTTP(w http.ResponseWriter, req *http.Request) {}
