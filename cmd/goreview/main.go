package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
)

var serverPort int

type (
	UserPayload struct {
		Name string `json:"login"`
	}

	RepoPayload struct {
		Name     string `json:"name"`
		FullName string `json:"full_name"`
	}

	PullRequestPayload struct {
		ID     int64  `json:"id"`
		URL    string `json:"url"`
		Number int64  `json:"number"`
		State  string `json:"state"`
	}

	PullRequestEventPayload struct {
		Action      string             `json:"action"`
		Repository  RepoPayload        `json:"repository"`
		User        UserPayload        `json:"sender"`
		PullRequest PullRequestPayload `json:"pull_request"`
	}
)

func init() {
	flag.IntVar(&serverPort, "port", 9999, "port to expose for server")
}

func egHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Print("TODO: replace me")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("FIX ME"))
}

func hookHandler(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	var p PullRequestEventPayload
	err = json.Unmarshal(b, &p)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	fmt.Printf("PULL REQUEST: %s\nBy: %s", p.PullRequest.URL, p.User.Name)
	w.WriteHeader(http.StatusAccepted)
	w.Write([]byte("OK"))
}

func main() {
	http.HandleFunc("/", egHandler)
	http.HandleFunc("/hooks", hookHandler)
	http.ListenAndServe(fmt.Sprintf(":%d", serverPort), nil)
}
