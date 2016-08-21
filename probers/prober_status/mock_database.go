package main

import (
    "fmt"
    "time"
)

var currentId int

var probes Probes

// Give us some seed data
func init() {
    RepoCreateProbe(Probe{Name: "Probe_1", Successful: true})
    RepoCreateProbe(Probe{Name: "Probe_2", Successful: false})
}

func RepoFindProbe(id int) Probe {
    for _, t := range probes {
        if t.Id == id {
            return t
        }
    }
    // return empty Probe if not found
    return Probe{}
}

func RepoCreateProbe(t Probe) Probe {
    currentId += 1
    t.Id = currentId
    now := time.Now()
    nanos := now.UnixNano()
    t.Timestamp = nanos
    probes = append(probes, t)
    return t
}

func RepoDestroyProbe(id int) error {
    for i, t := range probes {
        if t.Id == id {
            probes = append(probes[:i], probes[i+1:]...)
            return nil
        }
    }
    return fmt.Errorf("Could not find Probe with id of %d to delete", id)
}