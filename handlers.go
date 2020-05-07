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
	var greetings []types.GreetingRecord

	database.Get(greetings)
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
		newRecord, _ := database.Add(newGreeting)
		response := types.CreatedResponse{ID: newRecord.ID}

		encodeResponse(w, req, response, http.StatusCreated)
	}
}

func getCardHandlers(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodGet:
		getCards(w, req)
	case http.MethodPost:
		addCard(w, req)
	}
}

func getCards(w http.ResponseWriter, req *http.Request) {
	var cards []types.CardRecord
	database.Get(cards)
	encodeResponse(w, req, cards, http.StatusOK)
}

func addCard(w http.ResponseWriter, req *http.Request) {
	var newRecord types.Card
	err := decodeRequestBody(req, &newRecord)

	if err != nil {
		var res errorResponse = errorResponse{Error: "Internal Server Error"}
		encodedRes, _ := json.Marshal(res)

		w.WriteHeader(http.StatusInternalServerError)
		w.Write(encodedRes)
	} else {
		newID, _ := database.Add(newRecord)
		response := types.CreatedResponse{ID: newID.ID}

		encodeResponse(w, req, response, http.StatusCreated)
	}
}
