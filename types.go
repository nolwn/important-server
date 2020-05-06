package main

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
