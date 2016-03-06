package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

type WrapHTTPHandler struct {
	m http.Handler
}

func (h *WrapHTTPHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	lw := &loggedResponse{ResponseWriter: w, status: 200}
	start := time.Now()
	h.m.ServeHTTP(lw, r)
	elapsed := time.Since(start)
	log.SetPrefix("[Info]")
	log.Printf("[%s] %s - %d, time elapsed was: %dns.\n", 
		r.RemoteAddr, r.URL, lw.status, elapsed)
}

type loggedResponse struct {
	http.ResponseWriter
	status int
}

func (l *loggedResponse) WriteHeader(status int) {
	l.status = status
	l.ResponseWriter.WriteHeader(status)
}

func hello(w http.ResponseWriter, r *http.Request) {
	// The "/" pattern matches everything, so we need to check
    // that we're at the root here.
    if r.URL.Path != "/" {
            http.NotFound(w, r)
            return
    }
    fmt.Fprintf(w, "Welcome to the home page!")
}

func goodbye(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "goodbye")
}

func main() {
	http.HandleFunc("/", hello)
	http.HandleFunc("/goodbye", goodbye)
	http.Handle("/hello", http.RedirectHandler("/", http.StatusFound))
	log.Fatalln(http.ListenAndServe(":8080", &WrapHTTPHandler{http.DefaultServeMux}))
}