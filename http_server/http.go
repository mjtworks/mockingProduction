package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
  "flag"

  "github.com/prometheus/client_golang/prometheus"
)

// WrapHTTPHandler defines a new struct with a single named field called handler.
// handler is an http.Handler, which will come in handy when we want to wrap such Handlers
type WrapHTTPHandler struct {
	handler http.Handler
	stats map[string]RequestStats
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
  // run with "go run http.go -port=8090"
  portNumberFlag := flag.String("port", "8080", "the port number to run the http on")
  portNumber := ":"
  portNumber += *portNumberFlag
	stats := make(map[string]RequestStats)
	http.HandleFunc("/", rootHandler)
  http.Handle("/", prometheus.InstrumentHandler(
    "", http.FileServer(http.Dir("/usr/share/doc")),
  ))
	http.Handle("/redirect_me", http.RedirectHandler("/", http.StatusFound))
	log.Fatalln(http.ListenAndServe(portNumber, &WrapHTTPHandler{http.DefaultServeMux, stats}))
}
