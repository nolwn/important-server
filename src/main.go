package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type postGreeting struct {
	Greeting string `json:"greeting"`
}

var database = make([]string, 0, 10)

func main() {
	http.HandleFunc("/greeting", getGreetingHandlers)
	http.ListenAndServe(":3000", nil)
}

func getGreetingHandlers(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodGet:
		getGreetings(w, req)
	case http.MethodPost:
		addGreeting(w, req)
	}
}

func getGreetings(w http.ResponseWriter, req *http.Request) {
	bytes, err := json.Marshal(database)

	if err != nil {
		// TODO: Figure out how to set status code
		fmt.Fprintf(w, "Internal Server Error")
	} else {
		res := string(bytes)
		fmt.Fprintf(w, res)
	}
}

func addGreeting(w http.ResponseWriter, req *http.Request) {
	var newGreeting postGreeting
	err := json.NewDecoder(req.Body).Decode(&newGreeting)

	if err != nil {
		fmt.Fprintf(w, "Internal Server Error")
	}

	database = append(database, newGreeting.Greeting)
	fmt.Fprintf(w, "Created")
}
