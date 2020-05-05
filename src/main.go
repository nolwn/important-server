package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

var fieldCounter = 0

type postGreeting struct {
	Greeting string `json:"greeting"`
	Language string `json:"language"`
	IsFormal bool   `json:"isFormal"`
	NumWords int    `json:"numWords"`
}

type greetingRecord struct {
	ID       int    `json:"id"`
	Greeting string `json:"greeting"`
	Language string `json:"language"`
	IsFormal bool   `json:"isFormal"`
	NumWords int    `json:"numWords"`
}

type createdResponse struct {
	ID int `json:"id"`
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

func main() {
	http.HandleFunc("/greeting", getGreetingHandlers)
	err := http.ListenAndServe(":3000", nil)

	log.Fatal(err)
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
	encodeResponse(w, req, database, http.StatusOK)
}

func addGreeting(w http.ResponseWriter, req *http.Request) {
	var newGreeting postGreeting
	err := decodeRequestBody(req, &newGreeting)

	if err != nil {
		var res errorResponse = errorResponse{Error: "Internal Server Error"}
		encodedRes, _ := json.Marshal(res)

		w.WriteHeader(http.StatusInternalServerError)
		w.Write(encodedRes)
	} else {
		var newRecord greetingRecord
		makeNewRecord(&newGreeting, &newRecord)

		response := createdResponse{ID: newRecord.ID}
		database = append(database, newRecord)

		encodeResponse(w, req, response, http.StatusCreated)
	}
}

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
