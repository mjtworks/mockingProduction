package main

import (
	"fmt"
	"strconv"
	"log"
	"net/http"
	"time"
	"flag"

	"github.com/prometheus/client_golang/prometheus"
)

type WrapHTTPHandler struct {
	handler http.Handler
}

type LoggedResponse struct {
	http.ResponseWriter
	status int
}

var (
	httpResponsesTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: "mocking_production",
	        Subsystem: "http_server",
	        Name:      "http_responses_total",
	        Help:      "The count of http responses issued, classified by code and method.",
	    },
	    []string{"code", "method"},
	)

	httpResponseLatencies = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Namespace: "mocking_production",
	        Subsystem: "http_server",
	        Name:      "http_response_latencies",
	        Help:      "Distribution of http response latencies (ms), classified by code and method.",
	    },
	    []string{"code", "method"},
	)
)

func (loggedResponse *LoggedResponse) WriteHeader(status int) {
	loggedResponse.status = status
	loggedResponse.ResponseWriter.WriteHeader(status)
}

func (wrappedHandler *WrapHTTPHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	loggedWriter := &LoggedResponse{ResponseWriter: writer, status: 200}
	
	start := time.Now()
	wrappedHandler.handler.ServeHTTP(loggedWriter, request)
	elapsed := time.Since(start)
	msElapsed := elapsed / time.Millisecond

	status := strconv.Itoa(loggedWriter.status)
	httpResponsesTotal.WithLabelValues(status, request.Method).Inc()
	httpResponseLatencies.WithLabelValues(status, request.Method).Observe(float64(msElapsed))
	
	log.SetPrefix("[Info]")
	log.Printf("[%s] %s - %d, Method: %s, time elapsed was: %d(ms).\n",
		request.RemoteAddr, request.URL, loggedWriter.status, request.Method, msElapsed)
}

func rootHandler(writer http.ResponseWriter, request *http.Request) {
	if request.URL.Path != "/" {
		http.NotFound(writer, request)
		return
	}
	writer.WriteHeader(http.StatusOK)
	fmt.Fprintf(writer, "You've hit the home page.")
}

func init() {
    prometheus.MustRegister(httpResponsesTotal)
    prometheus.MustRegister(httpResponseLatencies)
}

func main() {
	// run with "go run http.go -port=8090"
	portNumberFlag := flag.String("port", "8080", "the port number to run the http on")
	// Once all flags are declared, call flag.Parse() to execute the command-line parsing.
	flag.Parse()
	portNumber := ":" + *portNumberFlag
	// Expose the registered metrics via the special prometheus metrics handler.
	http.Handle("/metrics", prometheus.Handler())

	http.HandleFunc("/", rootHandler)
	http.Handle("/redirect_me", http.RedirectHandler("/", http.StatusFound))
	log.Fatalln(http.ListenAndServe(portNumber, &WrapHTTPHandler{http.DefaultServeMux}))
}
