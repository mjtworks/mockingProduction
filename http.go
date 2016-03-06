package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

// WrapHTTPHandler defines a new struct with a single named field called m.
// m is an http.Handler, which will come in handy when we want to wrap such Handlers
type WrapHTTPHandler struct {
	m http.Handler
}

type loggedResponse struct {
	http.ResponseWriter
	status int
}

// ServeHTTP is a method with an WrapHTTPHandler as its receiver. 
// We use it to override the ServeHTTP methods of Handler class, so we can add things like logging 
// of the latency and status to it.
func (h *WrapHTTPHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	lw := &loggedResponse{ResponseWriter: w, status: 200}
	start := time.Now()
	h.m.ServeHTTP(lw, r)
	elapsed := time.Since(start)
	log.SetPrefix("[Info]")
	log.Printf("[%s] %s - %d, time elapsed was: %dns.\n", 
		r.RemoteAddr, r.URL, lw.status, elapsed)
}

func (l *loggedResponse) WriteHeader(status int) {
	l.status = status
	l.ResponseWriter.WriteHeader(status)
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	// The "/" pattern matches everything, so we need to check that we're at the root here.
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	fmt.Fprintf(w, "You've hit the home page.")
}

func main() {
	http.HandleFunc("/", rootHandler)
	http.Handle("/redirect_me", http.RedirectHandler("/", http.StatusFound))
	log.Fatalln(http.ListenAndServe(":8080", &WrapHTTPHandler{http.DefaultServeMux}))
}
