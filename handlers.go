package main

import (
	"encoding/json"
	"net/http"

	"github.com/nolwn/important-server/database"
	"github.com/nolwn/important-server/types"
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
	greetings := database.GetGreetings()
	encodeResponse(w, req, greetings, http.StatusOK)
}

func addGreeting(w http.ResponseWriter, req *http.Request) {
	var newGreeting types.PostGreeting
	err := decodeRequestBody(req, &newGreeting)

	if err != nil {
		var res errorResponse = errorResponse{Error: "Internal Server Error"}
		encodedRes, _ := json.Marshal(res)

		w.WriteHeader(http.StatusInternalServerError)
		w.Write(encodedRes)
	} else {
		newRecord := database.AddGreeting(newGreeting)
		response := types.CreatedResponse{ID: newRecord.ID}

		encodeResponse(w, req, response, http.StatusCreated)
	}
}
