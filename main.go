package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/greeting", getGreetingHandlers)
	http.HandleFunc("/card", getCardHandlers)
	err := http.ListenAndServe(":3000", nil)

	log.Fatal(err)
}

type errorResponse struct {
	Error string `json:"error"`
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
