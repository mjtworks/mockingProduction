package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
	"reflect"
)

// RequestStats tracks stats about the requests made to the server for later
// usage in monitoring and alerting.
type RequestStats struct {
	path string
	hitCount int
	errorCount int
	latency time.Duration
}

// WrapHTTPHandler defines a new struct with a single named field called handler.
// handler is an http.Handler, which will come in handy when we want to wrap such Handlers
type WrapHTTPHandler struct {
	handler http.Handler
	stats RequestStats
}

// LoggedResponse defines a struct that contains an http ResponseWriter and an
// integer HTTP status code. 
// This is used in the ServeHTTP method to provide a custom ResponseWriter that
// can send error codes, as opposed to the default 200 OK response.
type LoggedResponse struct {
	http.ResponseWriter
	status int
}

// ServeHTTP is a method with an WrapHTTPHandler as its receiver. 
// We use it to override the ServeHTTP methods of Handler class, so we can add things like logging 
// of the latency and status to it.
func (wrappedHandler *WrapHTTPHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	loggedWriter := &LoggedResponse{ResponseWriter: writer, status: 200}
	start := time.Now()
	wrappedHandler.handler.ServeHTTP(loggedWriter, request)
	elapsed := time.Since(start)
	fmt.Println(reflect.TypeOf(elapsed))
	log.SetPrefix("[Info]")
	log.Printf("[%s] %s - %d, time elapsed was: %dns.\n", 
		request.RemoteAddr, request.URL, loggedWriter.status, elapsed)
}

// WriteHeader overrides the WriteHeader provided by the ResponseWriter interface/
// It sends an HTTP response header with status code to the requester.
func (loggedResponse *LoggedResponse) WriteHeader(status int) {
	loggedResponse.status = status
	loggedResponse.ResponseWriter.WriteHeader(status)
}

// rootHandler takes care of requests for the root of the server, "/". It makes
// sure that the root is actually what is being requested, since DefaultServeMux
// matches anything under "/" as root. 
func rootHandler(writer http.ResponseWriter, request *http.Request) {
	// The "/" pattern matches everything, so we need to check that we're at the root here.
	if request.URL.Path != "/" {
		http.NotFound(writer, request)
		return
	}
	fmt.Fprintf(writer, "You've hit the home page.")
}

func main() {
	stats := new(RequestStats) // TODO: make this useful
	http.HandleFunc("/", rootHandler)
	http.Handle("/redirect_me", http.RedirectHandler("/", http.StatusFound))
	log.Fatalln(http.ListenAndServe(":8080", &WrapHTTPHandler{http.DefaultServeMux, *stats}))
}
