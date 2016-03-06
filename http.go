package main

import (
	"fmt"
	"net/http"
)


// handler takes in an http.Responsewriting and an http.Request .
// the ResponseWriter is an interface type and can contain either a copy of a pointer (reference-type) or a copy of a value (value-type).
// An http.ResponseWriter value assembles the HTTP server's response; by writing to it, we send data to the HTTP client.
// the Request is a reference, a data structure that represents the client HTTP request. 
// The trailing [1:] means "create a sub-slice of Path from the 1st character to the end," dropping the leading "/" from the path name.
func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hi there, I love %s! The method is %s.\n", r.URL.Path[1:], r.Method)
    // send POST requests with $ curl -X POST localhost:8080/something
}

func main() {
    http.HandleFunc("/", handler) // handle all requests to the web root ("/") with handler function
    http.ListenAndServe(":8080", nil) // listen on port 8080 on any interface (":8080"), blocking until terminated
}