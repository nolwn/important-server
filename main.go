package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/greeting", getGreetingHandlers)
	err := http.ListenAndServe(":3000", nil)

	log.Fatal(err)
}

func makeNewRecord(pg *postGreeting, gr *greetingRecord) {
	gr.ID = fieldCounter
	gr.Greeting = pg.Greeting
	gr.Language = pg.Language
	gr.IsFormal = pg.IsFormal
	gr.NumWords = pg.NumWords

	fieldCounter++
}

type errorResponse struct {
	Error string `json:"error"`
}

var database = make([]greetingRecord, 0, 10)

func decodeRequestBody(req *http.Request, i interface{}) error {
	err := json.NewDecoder(req.Body).Decode(&i)
	return err
}

func encodeResponse(w http.ResponseWriter, req *http.Request, i interface{}, code int) ([]byte, error) {
	bytes, err := json.Marshal(i)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Internal Server Error")
	} else {
		w.WriteHeader(code)
		w.Write(bytes)
	}

	return bytes, err
}
