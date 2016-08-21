package main

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"time"
	"io"
	"io/ioutil"
	"github.com/gorilla/mux"
)

type Probe struct {
	Id			int			`json: "id"`
	Name 		string 		`json:"name"`
	Successful  bool 		`json:"succesful"`
	Timestamp 	int64	 	`json:"timestamp"`
}

type Probes []Probe

func Logger(r *http.Request, name string) {
        start := time.Now()
        log.Printf(
            "%s\t%s\t%s\t%s",
            r.Method,
            r.RequestURI,
            name,
            time.Since(start) / time.Millisecond,
        )
}

func main() {
    router := mux.NewRouter().StrictSlash(true)
    router.HandleFunc("/", ViewProbes)
    router.HandleFunc("/create", CreateProbe)
    router.HandleFunc("/view", ViewProbes)

    log.Fatal(http.ListenAndServe(":8080", router))
}

func CreateProbe(w http.ResponseWriter, r *http.Request) {
	// example json curl request to create probe event:
	// curl -H "Content-Type: application/json" -d '{"name":"probe3","succesful":true}' http://localhost:8080/create
	Logger(r, "create probe")
	var probe Probe
	// open the body of the request, but limit what we read in
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
    if err != nil {
        panic(err)
    }
    if err := r.Body.Close(); err != nil {
        panic(err)
    }
    if err := json.Unmarshal(body, &probe); err != nil {
        w.Header().Set("Content-Type", "application/json; charset=UTF-8")
        w.WriteHeader(422) // unprocessable entity
        if err := json.NewEncoder(w).Encode(err); err != nil {
            panic(err)
        }
    }

    t := RepoCreateProbe(probe) // create and store the probe result
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusCreated) // return 201
    if err := json.NewEncoder(w).Encode(t); err != nil {
        panic(err)
    }
}

func ViewProbes(w http.ResponseWriter, r *http.Request) {
	// TODO(pheven): update this so that it displays the probes in an html table
	Logger(r, "view probes")
	log.Printf("Number of probes: %d", currentId)
	for i := 0; i < currentId + 1; i++ {
		t := RepoFindProbe(i)
		json.NewEncoder(w).Encode(t)
	}
	t, _ := template.ParseFiles("response.html")
	t.Execute(w, r)
}
