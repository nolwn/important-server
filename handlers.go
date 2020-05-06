package main

import (
	"encoding/json"
	"net/http"
)

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
